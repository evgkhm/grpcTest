package usecase

import (
	"context"
	"github.com/PuerkitoBio/goquery"
	"github.com/pkg/errors"
	"grpcTest/internal/entity"
	"io"
	"log/slog"
	"net/http"
	"strconv"
	"strings"
)

var (
	ErrItnNotFound = errors.New("Itn not found")
)

type ItnUseCase struct {
	logger *slog.Logger
}

func NewItnUseCase(logger *slog.Logger) *ItnUseCase {
	return &ItnUseCase{
		logger: logger,
	}
}

// GetItnInfo retrieves the information of an ITN.
//
// It takes the following parameters:
// - ctx: the context.Context object for handling deadlines and cancellations.
// - inputItn: the ID of the ITN to retrieve information for.
//
// It returns a pointer to a DTO object and an error.
func (i *ItnUseCase) GetItnInfo(ctx context.Context, inputItn uint64) (*entity.DTO, error) {
	doc, err := makeReq(ctx, inputItn)
	if err != nil {
		i.logger.Error("GetItnInfo - makeReq", err)
		return nil, err
	}

	err = checkDoc(doc)
	if err != nil {
		i.logger.Error("GetItnInfo - checkDoc", err)
		return nil, err
	}

	dto, parseErr := parseDoc(doc)
	if parseErr != nil {
		i.logger.Error("GetItnInfo - parseDoc", err)
		return nil, parseErr
	}

	return dto, nil
}

// checkDoc checks the given goquery.Document for the presence of the text "результаты поиска" in the title tag.
//
// It takes a pointer to a goquery.Document as a parameter.
// It does not return any value.
func checkDoc(doc *goquery.Document) error {
	doc.Find("title").EachWithBreak(func(i int, s *goquery.Selection) bool {
		if strings.Contains(s.Text(), "результаты поиска") == true {
			return false
		}
		return true
	})
	return nil
}

// makeReq retrieves a goquery.Document from a URL based on the input ITN.
//
// It takes a context.Context and an unsigned 64-bit integer as its parameters.
// It returns a pointer to a goquery.Document and an error.
func makeReq(ctx context.Context, inputItn uint64) (*goquery.Document, error) {
	url := "https://www.rusprofile.ru/search?query="
	itnStr := strconv.Itoa(int(inputItn))

	req, err := http.Get(url + itnStr)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			return
		}
	}(req.Body)

	doc, err := goquery.NewDocumentFromReader(req.Body)
	if err != nil {
		return nil, err
	}

	return doc, err
}

// parseDoc parses the given goquery.Document and returns a DTO and an error.
//
// It takes a *goquery.Document as a parameter and searches for specific elements within it. It populates a DTO struct with the found values and returns it along with an error.
// The function returns *entity.DTO and an error.
func parseDoc(doc *goquery.Document) (*entity.DTO, error) {
	var companyDTO entity.DTO
	doc.Find("span").EachWithBreak(func(index int, item *goquery.Selection) bool {
		val, ok := item.Attr("id")
		if ok && val == "clip_inn" {
			itemStr := item.Text()
			itemUint, err := strconv.ParseUint(itemStr, 10, 64)
			if err != nil {
				return false
			}
			companyDTO.Itn = itemUint
			return false
		}
		return true
	})

	if companyDTO.Itn == 0 {
		return nil, ErrItnNotFound
	}

	doc.Find("span").EachWithBreak(func(index int, item *goquery.Selection) bool {
		val, ok := item.Attr("id")
		if ok && val == "clip_kpp" {
			itemStr := item.Text()
			itemUint, err := strconv.ParseUint(itemStr, 10, 64)
			if err != nil {
				return false
			}
			companyDTO.Kpp = itemUint
			return false

		}
		return true
	})

	doc.Find("h1").EachWithBreak(func(index int, item *goquery.Selection) bool {
		val, ok := item.Attr("itemprop")
		if ok && val == "name" {
			companyDTO.NameCompany = item.Text()
			companyDTO.NameCompany = strings.Replace(companyDTO.NameCompany, `"`, "", -1)
			return false
		}
		return true
	})

	doc.Find("span").EachWithBreak(func(index int, item *goquery.Selection) bool {
		val, ok := item.Attr("class")
		if ok && val == "company-info__text" {
			companyDTO.NameLeader = item.Text()
			return false

		}
		return true
	})

	return &companyDTO, nil
}

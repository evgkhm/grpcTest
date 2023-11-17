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

func (i *ItnUseCase) GetItnInfo(ctx context.Context, dto *entity.DTO) error {
	doc, err := makeReq(ctx, dto)
	if err != nil {
		i.logger.Error("GetItnInfo - makeReq", err)
		return err
	}

	err = checkDoc(doc, dto)
	if err != nil {
		i.logger.Error("GetItnInfo - checkDoc", err)
		return err
	}

	err = parseDoc(doc, dto)
	if err != nil {
		i.logger.Error("GetItnInfo - parseDoc", err)
		return err
	}

	return nil
}

func checkDoc(doc *goquery.Document, dto *entity.DTO) error {
	doc.Find("title").EachWithBreak(func(i int, s *goquery.Selection) bool {
		if strings.Contains(s.Text(), "результаты поиска") == true {
			return false
		}
		return true
	})
	return nil
}

func makeReq(ctx context.Context, dto *entity.DTO) (*goquery.Document, error) {
	url := "https://www.rusprofile.ru/search?query="
	itnStr := strconv.Itoa(int(dto.Itn))

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

func parseDoc(doc *goquery.Document, companyDTO *entity.DTO) error {
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
		return ErrItnNotFound
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

	return nil
}

func NewItnUseCase(logger *slog.Logger) *ItnUseCase {
	return &ItnUseCase{
		logger: logger,
	}
}

package linkservice

import (
	"ginexamples"
	"ginexamples/pkg/auth"
	"regexp"

	"github.com/pkg/errors"
)

const urlRe = `^(?:http(s)?:\/\/)?[\w.-]+(?:\.[\w\.-]+)+[\w\-\._~:/?#[\]@!\$&'\(\)\*\+,;=.]+$`

type LinkService struct {
	r ginexamples.LinkRepository
}

func New(linkRepository ginexamples.LinkRepository) *LinkService {
	return &LinkService{
		r: linkRepository,
	}
}

func (lS *LinkService) CreateLink(link *ginexamples.Link, url string) (*ginexamples.Link, error) {
	_, err := lS.r.FindByUserName(link.UserName)
	if err == nil {
		return &ginexamples.Link{}, errors.New("username already exists")
	}

	//add by andy
	ok, err := regexp.MatchString(urlRe, url)
	if !ok || err != nil {
		return &ginexamples.Link{}, errors.Wrap(err, "error url is illgal")
	}

	err = lS.r.Store(link)
	if err != nil {
		return &ginexamples.Link{}, errors.Wrap(err, "error storing link")
	}
	return link, nil
}

func (lS *LinkService) GetLink(id string) (*ginexamples.Link, error) {
	return lS.r.Find(id)
}

func (lS *LinkService) UpdateLink(link *ginexamples.Link) error {
	return lS.r.Update(link)
}

func (lS *LinkService) DeleteLink(id string) (*ginexamples.Link, error) {
	return lS.r.Delete(id)
}

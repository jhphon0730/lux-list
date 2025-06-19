package service

import "lux-list/internal/repository"

// TagService는 태그 관련 메서드를 정의하는 인터페이스
type TagService interface{}

// tagService는 TagService 인터페이스를 구현하는 구조체
type tagService struct {
	tagRepository repository.TagRepository
}

// NewTagService는 TagService의 인스턴스를 생성하는 함수
func NewTagService(tagRepository repository.TagRepository) TagService {
	return &tagService{
		tagRepository: tagRepository,
	}
}

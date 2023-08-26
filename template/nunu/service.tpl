package module

import (
	"context"

	"{{ .ProjectName }}/service/internal/model/bo"
	"{{ .ProjectName }}/service/internal/model/po"

	"github.com/jinzhu/copier"
)

type I{{ .FileName }}Module interface {
	Get{{ .FileName }}ByID(ctx context.Context, cond *bo.Query{{ .FileName }}Cond) (*bo.{{ .FileName }}, error)
	Get{{ .FileName }}(ctx context.Context, cond *bo.Query{{ .FileName }}Cond) ([]*bo.{{ .FileName }}, *bo.Page, error)
	Create{{ .FileName }}(ctx context.Context, cond *bo.Create{{ .FileName }}Cond) (*bo.{{ .FileName }}, error)
	Update{{ .FileName }}(ctx context.Context, cond *bo.Update{{ .FileName }}Cond) (*bo.{{ .FileName }}, error)
	Delete{{ .FileName }}(ctx context.Context, cond *bo.Delete{{ .FileName }}Cond) error
}

func New{{ .FileName }}Module(in digIn) I{{ .FileName }}Module {
	return &{{ .FileNameTitleLower }}Module{in: in}
}

type {{ .FileNameTitleLower }}Module struct {
	in digIn
}

func (m *{{ .FileNameTitleLower }}Module) Get{{ .FileName }}ByID(ctx context.Context, cond *bo.Query{{ .FileName }}Cond) (*bo.{{ .FileName }}, error) {
	poCond := &po.Query{{ .FileName }}Cond{}
	if err := copier.Copy(poCond, cond); err != nil {
		return nil, err
	}

	poResult, err := m.in.Repository.{{ .FileName }}Repo.Get{{ .FileName }}ByID(ctx, poCond)
	if err != nil {
		return nil, err
	}
	result := &bo.{{ .FileName }}{}

	if err := copier.Copy(result, poResult); err != nil {
		return nil, err
	}

	return result, nil
}

func (m *{{ .FileNameTitleLower }}Module) Get{{ .FileName }}(ctx context.Context, cond *bo.Query{{ .FileName }}Cond) ([]*bo.{{ .FileName }}, *bo.Page, error) {
	poCond := &po.Query{{ .FileName }}Cond{}
	if err := copier.Copy(poCond, cond); err != nil {
		return nil, nil, err
	}

	poResult, poPager, err := m.in.Repository.{{ .FileName }}Repo.Get{{ .FileName }}(ctx, poCond)
	if err != nil {
		return nil, nil, err
	}
	result := make([]*bo.{{ .FileName }}, len(poResult))
	if err := copier.Copy(&result, poResult); err != nil {
		return nil, nil, err
	}
	page := &bo.Page{}
	if err := copier.Copy(page, poPager); err != nil {
		return nil, nil, err
	}

	return result, page, nil
}

func (m *{{ .FileNameTitleLower }}Module) Create{{ .FileName }}(ctx context.Context, cond *bo.Create{{ .FileName }}Cond) (*bo.{{ .FileName }}, error) {
	{{ .FileNameTitleLower }} := &po.{{ .FileName }}{}
	if err := copier.Copy({{ .FileNameTitleLower }}, cond); err != nil {
		return nil, err
	}

	err := m.in.Repository.{{ .FileName }}Repo.Create{{ .FileName }}(ctx, {{ .FileNameTitleLower }})
	if err != nil {
		return nil, err
	}

	result := &bo.{{ .FileName }}{}
	if err := copier.Copy(result, {{ .FileNameTitleLower }}); err != nil {
		return nil, err
	}
	return result, nil
}

func (m *{{ .FileNameTitleLower }}Module) Update{{ .FileName }}(ctx context.Context, cond *bo.Update{{ .FileName }}Cond) (*bo.{{ .FileName }}, error) {
	{{ .FileNameTitleLower }} := &po.{{ .FileName }}{}
	if err := copier.Copy({{ .FileNameTitleLower }}, cond); err != nil {
		return nil, err
	}

	err := m.in.Repository.{{ .FileName }}Repo.Update{{ .FileName }}(ctx, {{ .FileNameTitleLower }})
	if err != nil {
		return nil, err
	}

	result := &bo.{{ .FileName }}{}
	if err := copier.Copy(result, {{ .FileNameTitleLower }}); err != nil {
		return nil, err
	}
	return result, nil
}

func (m *{{ .FileNameTitleLower }}Module) Delete{{ .FileName }}(ctx context.Context, cond *bo.Delete{{ .FileName }}Cond) error {
	{{ .FileNameTitleLower }} := &po.{{ .FileName }}{}
	if err := copier.Copy({{ .FileNameTitleLower }}, cond); err != nil {
		return err
	}
	err := m.in.Repository.{{ .FileName }}Repo.Delete{{ .FileName }}(ctx, {{ .FileNameTitleLower }})
	if err != nil {
		return err
	}
	return nil
}

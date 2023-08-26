package repository

import (
	"context"

	"{{ .ProjectName }}/service/internal/model/po"
	"{{ .ProjectName }}/service/internal/util/ctxs"
)

type I{{ .FileName }}Repository interface {
	Get{{ .FileName }}ByID(ctx context.Context, cond *po.Query{{ .FileName }}Cond) (*po.{{ .FileName }}, error)
	Get{{ .FileName }}(ctx context.Context, cond *po.Query{{ .FileName }}Cond) ([]*po.{{ .FileName }}, *po.Page, error)
	Create{{ .FileName }}(ctx context.Context, cond *po.{{ .FileName }}) error
	Update{{ .FileName }}(ctx context.Context, cond *po.{{ .FileName }}) error
	Delete{{ .FileName }}(ctx context.Context, cond *po.{{ .FileName }}) error
}

type {{ .FileNameTitleLower }}Repository struct {
	in digIn
}

func New{{ .FileName }}Repo(in digIn) I{{ .FileName }}Repository {
	return &{{ .FileNameTitleLower }}Repository{
		in: in,
	}
}

func (m *{{ .FileNameTitleLower }}Repository) Get{{ .FileName }}ByID(ctx context.Context, cond *po.Query{{ .FileName }}Cond) (*po.{{ .FileName }}, error) {
	{{ .FileNameTitleLower }} := &po.{{ .FileName }}{}
	if err := m.in.GetDB(ctx).Where(cond).Find({{ .FileNameTitleLower }}, cond).Error; err != nil {
		return nil, err
	}
	return {{ .FileNameTitleLower }}, nil
}

func (m *{{ .FileNameTitleLower }}Repository) Get{{ .FileName }}(ctx context.Context, cond *po.Query{{ .FileName }}Cond) ([]*po.{{ .FileName }}, *po.Page, error) {
	{{ .FileNameTitleLower }} := make([]*po.{{ .FileName }}, 0)
	c := ctxs.SetValue(ctx, "pager", cond.GetPager())
	if err := m.in.GetDB(c).Model(po.{{ .FileName }}{}).Find(&{{ .FileNameTitleLower }}, cond).Error; err != nil {
		return nil, nil, err
	}
	return {{ .FileNameTitleLower }}, cond.GetPager(), nil
}

func (m *{{ .FileNameTitleLower }}Repository) Create{{ .FileName }}(ctx context.Context, {{ .FileNameTitleLower }} *po.{{ .FileName }}) error {
	if err := m.in.GetDB(ctx).Create({{ .FileNameTitleLower }}).Error; err != nil {
		return err
	}
	return nil
}

func (m *{{ .FileNameTitleLower }}Repository) Update{{ .FileName }}(ctx context.Context, {{ .FileNameTitleLower }} *po.{{ .FileName }}) error {
	if err := m.in.GetDB(ctx).Updates({{ .FileNameTitleLower }}).Error; err != nil {
		return err
	}
	return nil
}

func (m *{{ .FileNameTitleLower }}Repository) Delete{{ .FileName }}(ctx context.Context, {{ .FileNameTitleLower }} *po.{{ .FileName }}) error {
	if err := m.in.GetDB(ctx).Delete(&po.{{ .FileName }}{}, {{ .FileNameTitleLower }}).Error; err != nil {
		return err
	}
	return nil
}

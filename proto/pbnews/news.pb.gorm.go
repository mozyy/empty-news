// Code generated by protoc-gen-gorm. DO NOT EDIT.
// versions:
// 	protoc-gen-gorm v2.0.0
// 	protoc        v3.19.3
// source: proto/news/news.proto

package pbnews

import (
	context "context"
	_ "github.com/mozyy/protoc-gen-gorm/options"
	types "github.com/mozyy/protoc-gen-gorm/types"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	time "time"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type NewsItemGORM struct {
	ID        uint32 `gorm:"primaryKey"`
	Link      string
	Image     string //
	Title     string
	Time      string
	View      int32
	Comment   int32
	Type      string // 1: 业界, 2: 科学, 3: 影视, 4: 游戏
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *types.DeletedAt `gorm:"index"`
}

// TableName overrides the default tablename generated by GORM
func (NewsItemGORM) TableName() string {
	return "news_items"
}

// ToPB  converts the fields of this object to PB object
func (m *NewsItemGORM) ToPB(ctx context.Context) *NewsItem {
	to := NewsItem{}
	to.ID = m.ID
	to.Link = m.Link
	to.Image = m.Image
	to.Title = m.Title
	to.Time = m.Time
	to.View = m.View
	to.Comment = m.Comment
	to.Type = m.Type
	if m.CreatedAt != nil {
		to.CreatedAt = timestamppb.New(*m.CreatedAt)
	}
	if m.UpdatedAt != nil {
		to.UpdatedAt = timestamppb.New(*m.UpdatedAt)
	}
	to.DeletedAt = m.DeletedAt
	return &to
}

// ToORM  converts the fields of this object to PB object
func (m *NewsItem) ToORM(ctx context.Context) *NewsItemGORM {
	to := NewsItemGORM{}
	to.ID = m.GetID()
	to.Link = m.GetLink()
	to.Image = m.GetImage()
	to.Title = m.GetTitle()
	to.Time = m.GetTime()
	to.View = m.GetView()
	to.Comment = m.GetComment()
	to.Type = m.GetType()
	if m.GetCreatedAt() != nil {
		CreatedAt := m.GetCreatedAt().AsTime()
		to.CreatedAt = &CreatedAt
	}
	if m.GetUpdatedAt() != nil {
		UpdatedAt := m.GetUpdatedAt().AsTime()
		to.UpdatedAt = &UpdatedAt
	}
	to.DeletedAt = m.GetDeletedAt()
	return &to
}

type DetailResponseGORM struct {
	ID        uint32 `gorm:"primaryKey"`
	Title     string
	From      string
	Time      string
	Summary   string
	Content   []*DetailResponse_DetailContentGORM
	No        string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *types.DeletedAt `gorm:"index"`
}

// TableName overrides the default tablename generated by GORM
func (DetailResponseGORM) TableName() string {
	return "news_details"
}

// ToPB  converts the fields of this object to PB object
func (m *DetailResponseGORM) ToPB(ctx context.Context) *DetailResponse {
	to := DetailResponse{}
	to.ID = m.ID
	to.Title = m.Title
	to.From = m.From
	to.Time = m.Time
	to.Summary = m.Summary
	for _, Content := range m.Content {
		to.Content = append(to.Content, Content.ToPB(ctx))
	}
	to.No = m.No
	if m.CreatedAt != nil {
		to.CreatedAt = timestamppb.New(*m.CreatedAt)
	}
	if m.UpdatedAt != nil {
		to.UpdatedAt = timestamppb.New(*m.UpdatedAt)
	}
	to.DeletedAt = m.DeletedAt
	return &to
}

// ToORM  converts the fields of this object to PB object
func (m *DetailResponse) ToORM(ctx context.Context) *DetailResponseGORM {
	to := DetailResponseGORM{}
	to.ID = m.GetID()
	to.Title = m.GetTitle()
	to.From = m.GetFrom()
	to.Time = m.GetTime()
	to.Summary = m.GetSummary()
	for _, Content := range m.GetContent() {
		to.Content = append(to.Content, Content.ToORM(ctx))
	}
	to.No = m.GetNo()
	if m.GetCreatedAt() != nil {
		CreatedAt := m.GetCreatedAt().AsTime()
		to.CreatedAt = &CreatedAt
	}
	if m.GetUpdatedAt() != nil {
		UpdatedAt := m.GetUpdatedAt().AsTime()
		to.UpdatedAt = &UpdatedAt
	}
	to.DeletedAt = m.GetDeletedAt()
	return &to
}

type DetailResponse_DetailContentGORM struct {
	ID           uint32 `gorm:"primaryKey"`
	Type         int32
	Content      string
	NewsDetailID uint32
	CreatedAt    *time.Time
	UpdatedAt    *time.Time
	DeletedAt    *types.DeletedAt `gorm:"index"`
}

// TableName overrides the default tablename generated by GORM
func (DetailResponse_DetailContentGORM) TableName() string {
	return "detail_response_detail_contents"
}

// ToPB  converts the fields of this object to PB object
func (m *DetailResponse_DetailContentGORM) ToPB(ctx context.Context) *DetailResponse_DetailContent {
	to := DetailResponse_DetailContent{}
	to.ID = m.ID
	to.Type = m.Type
	to.Content = m.Content
	to.NewsDetailID = m.NewsDetailID
	if m.CreatedAt != nil {
		to.CreatedAt = timestamppb.New(*m.CreatedAt)
	}
	if m.UpdatedAt != nil {
		to.UpdatedAt = timestamppb.New(*m.UpdatedAt)
	}
	to.DeletedAt = m.DeletedAt
	return &to
}

// ToORM  converts the fields of this object to PB object
func (m *DetailResponse_DetailContent) ToORM(ctx context.Context) *DetailResponse_DetailContentGORM {
	to := DetailResponse_DetailContentGORM{}
	to.ID = m.GetID()
	to.Type = m.GetType()
	to.Content = m.GetContent()
	to.NewsDetailID = m.GetNewsDetailID()
	if m.GetCreatedAt() != nil {
		CreatedAt := m.GetCreatedAt().AsTime()
		to.CreatedAt = &CreatedAt
	}
	if m.GetUpdatedAt() != nil {
		UpdatedAt := m.GetUpdatedAt().AsTime()
		to.UpdatedAt = &UpdatedAt
	}
	to.DeletedAt = m.GetDeletedAt()
	return &to
}

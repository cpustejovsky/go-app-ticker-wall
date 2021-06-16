package models

// UpdateType is a constant used to define the contents of an update.
type UpdateType int32

const (
	UpdateTypeUnknown      UpdateType = 0
	UpdateTypeCluster      UpdateType = 1
	UpdateTypeTicker       UpdateType = 2
	UpdateTypeAnnouncement UpdateType = 3
	UpdateTypePrice        UpdateType = 4
)

// AnnouncementType is used to signify the type of announcement / alert. Different announcement types behave differently.
type AnnouncementType int32

const (
	AnnouncementTypeInfo    AnnouncementType = 0
	AnnouncementTypeDanger  AnnouncementType = 1
	AnnouncementTypeSuccess AnnouncementType = 2
)

// AnnouncementAnimation are the different animation options available for an announcement.
type AnnouncementAnimation int32

const (
	AnnouncementAnimationElastic AnnouncementAnimation = 0
	AnnouncementAnimationBounce  AnnouncementAnimation = 1
	AnnouncementAnimationEase    AnnouncementAnimation = 2
	AnnouncementAnimationBack    AnnouncementAnimation = 3
)

package internal

import "github.com/hvxahv/hvx/cockroach"

type Editor interface {
	// EditTitle is a method for editing the title of an article.
	EditTitle(title string) *Articles

	// EditSummary is a method for editing the summary of an article.
	EditSummary(summary string) *Articles

	// EditArticle is a method for editing the article of an article.
	EditArticle(article string) *Articles

	// EditTags is a method for editing the tags of an article.
	EditTags(tags []string) *Articles

	// EditAttachmentType is a method for editing the attachment article.
	EditAttachmentType(attachmentType string) *Articles

	// EditAttachments is a method for editing the attachments of an article.
	EditAttachments(attachments []string) *Articles

	// EditNSFW is a method for editing the nsfw of an article.
	EditNSFW(nsfw bool) *Articles

	// EditVisibility is a method for editing the visibility of an article.
	EditVisibility(visibility uint) *Articles
}

func (a *Articles) EditTitle(title string) *Articles {
	a.Title = title
	return a
}

func (a *Articles) EditSummary(summary string) *Articles {
	a.Summary = summary
	return a
}

func (a *Articles) EditArticle(article string) *Articles {
	a.Article = article
	return a
}

func (a *Articles) EditTags(tags []string) *Articles {
	a.Tags = tags
	return a
}

func (a *Articles) EditAttachments(attachments []string) *Articles {
	a.Attachments = attachments
	return a
}

func (a *Articles) EditNSFW(nsfw bool) *Articles {
	a.NSFW = nsfw
	return a
}

func (a *Articles) EditVisibility(visibility int64) *Articles {
	a.Visibility = visibility
	return a
}

func (a *Articles) Edit() error {
	db := cockroach.GetDB()
	if err := db.Debug().
		Table(ArticleTable).
		Where("id = ? AND actor_id = ?", a.ID, a.ActorId).
		Updates(a).
		Error; err != nil {
		return nil
	}
	return nil
}

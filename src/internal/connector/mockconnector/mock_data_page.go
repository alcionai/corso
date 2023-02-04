package mockconnector

// GetMockPage returns bytes for models.SitePageable object
// Title string changes of fields: name and title
func GetMockPage(title string) []byte {
	fileName := title + ".aspx"

	// Create Test Page
	//nolint:lll
	byteArray := []byte("{\"name\":\"" + fileName + "\",\"title\":\"" + title + "\",\"pageLayout\":\"article\",\"showComments\":true," +
		"\"showRecommendedPages\":false,\"titleArea\":{\"enableGradientEffect\":true,\"imageWebUrl\":\"/_LAYOUTS/IMAGES/VISUALTEMPLATETITLEIMAGE.JPG\"," +
		"\"layout\":\"colorBlock\",\"showAuthor\":true,\"showPublishedDate\":false,\"showTextBlockAboveTitle\":false,\"textAboveTitle\":\"TEXTABOVETITLE\"," +
		"\"textAlignment\":\"left\",\"imageSourceType\":2,\"title\":\"sample1\"}," +
		"\"canvasLayout\":{\"horizontalSections\":[{\"layout\":\"oneThirdRightColumn\",\"id\":\"1\",\"emphasis\":\"none\",\"columns\":[{\"id\":\"1\",\"width\":8," +
		"\"webparts\":[{\"id\":\"6f9230af-2a98-4952-b205-9ede4f9ef548\",\"innerHtml\":\"<p><b>Hello!</b></p>\"}]},{\"id\":\"2\",\"width\":4," +
		"\"webparts\":[{\"id\":\"73d07dde-3474-4545-badb-f28ba239e0e1\",\"webPartType\":\"d1d91016-032f-456d-98a4-721247c305e8\",\"data\":{\"dataVersion\":\"1.9\"," +
		"\"description\":\"Showanimageonyourpage\",\"title\":\"Image\",\"properties\":{\"imageSourceType\":2,\"altText\":\"\",\"overlayText\":\"\"," +
		"\"siteid\":\"0264cabe-6b92-450a-b162-b0c3d54fe5e8\",\"webid\":\"f3989670-cd37-4514-8ccb-0f7c2cbe5314\",\"listid\":\"bdb41041-eb06-474e-ac29-87093386bb14\"," +
		"\"uniqueid\":\"d9f94b40-78ba-48d0-a39f-3cb23c2fe7eb\",\"imgWidth\":4288,\"imgHeight\":2848,\"fixAspectRatio\":false,\"captionText\":\"\",\"alignment\":\"Center\"}," +
		"\"serverProcessedContent\":{\"imageSources\":[{\"key\":\"imageSource\",\"value\":\"/_LAYOUTS/IMAGES/VISUALTEMPLATEIMAGE1.JPG\"}]," +
		"\"customMetadata\":[{\"key\":\"imageSource\",\"value\":{\"siteid\":\"0264cabe-6b92-450a-b162-b0c3d54fe5e8\",\"webid\":\"f3989670-cd37-4514-8ccb-0f7c2cbe5314\"," +
		"\"listid\":\"bdb41041-eb06-474e-ac29-87093386bb14\",\"uniqueid\":\"d9f94b40-78ba-48d0-a39f-3cb23c2fe7eb\",\"width\":\"4288\",\"height\":\"2848\"}}]}}}]}]}]}}")

	return byteArray
}

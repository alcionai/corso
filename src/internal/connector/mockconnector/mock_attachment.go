package mockconnector

func GetMockAttachmentReference() []byte {
	//nolint:lll
	attachment := "\"@odata.type\":\"#microsoft.graph.referenceAttachment\",\"id\":\"AAMkAGQ1NzViZTdhLTEwMTMtNGJjNi05YWI2LTg4NWRlZDA2Y2UxOABGAAAAAAAPvVwUramXT7jlSGpVU8_7BwB8wYc0thTTTYl3RpEYIUq_AAAAAAEJAAB8wYc0thTTTYl3RpEYIUq_AAB0tLkCAAABEgAQALqePWc7rIJHp_oQ7U_XyOI=\"," +
		"\"lastModifiedDateTime\":\"2022-09-30T19:33:26Z\",\"name\":\"sample.txt\",\"contentType\":\"text/plain\",\"size\":1042,\"isInline\":true}"

	return []byte(attachment)
}

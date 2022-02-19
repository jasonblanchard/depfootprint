pushweb:
	aws s3 sync ./web/depfootprint/out s3://depfootprint-static-assets --acl public-read
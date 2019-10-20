package pixivgogo

//go:generate sh -c "rm -f mock/*"
//go:generate mockgen -destination=mock/client_mock.go -package=mock -source ./client.go -mock_names urlValuesEncoder=URLValuesEncoder,urlValuesDecoder=URLValuesDecoder

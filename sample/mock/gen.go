package mock

//go:generate moq -out sample.go -pkg mock -stub -skip-ensure -with-resets -rm -fmt gofmt ../ Sample

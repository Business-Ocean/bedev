cobra-setup:
		cobra-cli init --config ".cobra.yaml"

build:
	rm -rf bedev &&	go build -o bedev
		
run:
	./bedev
tidy:
	go mod tidy

ppr:
	git push origin faraji && gh pr create --web -B dev
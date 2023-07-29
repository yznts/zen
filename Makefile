
# Serve docs
doc:
	(sleep 1 && open http://localhost:8000/pkg/go.kyoto.codes/zen/v3) &
	godoc -http=:8000

import http from "k6/http"
import { sleep } from "k6"

export const options = {
	stages: [
		{ duration: "5m", target: 1000 },
		{ duration: "10m", target: 1000 },
		{ duration: "5m", target: 0 },
	]
}
// TODO all servers.. gulp
export default function() {
	http.get("http://localhost:9000/healthcheckhr")
	sleep(1)
}


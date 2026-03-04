package ui

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func StartServer(port string) {
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(indexHTML))
	})

	log.Printf("[UI] Starting mesh visualizer on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("UI server failed: %v", err)
	}
}

const indexHTML = `
<!DOCTYPE html>
<html>
<head>
    <title>Lumina-Mesh Visualizer</title>
    <script src="https://d3js.org/d3.v7.min.js"></script>
    <style>
        body { font-family: sans-serif; background: #111; color: #fff; }
        .node { stroke: #fff; stroke-width: 1.5px; fill: #69b3a2; }
        .link { stroke: #999; stroke-opacity: 0.6; }
    </style>
</head>
<body>
    <h1>Lumina-Mesh Real-time Visualizer</h1>
    <div id="mesh"></div>
    <script>
        const width = 800, height = 600;
        const svg = d3.select("#mesh").append("svg").attr("width", width).attr("height", height);
        
        // Dynamic nodes simulation
        let nodes = [{id: "Node-1"}, {id: "Node-2"}, {id: "Node-3"}];
        let links = [{source: "Node-1", target: "Node-2"}, {source: "Node-2", target: "Node-3"}];

        const simulation = d3.forceSimulation(nodes)
            .force("link", d3.forceLink(links).id(d => d.id))
            .force("charge", d3.forceManyBody())
            .force("center", d3.forceCenter(width / 2, height / 2));

        const link = svg.append("g").selectAll("line").data(links).join("line").attr("class", "link");
        const node = svg.append("g").selectAll("circle").data(nodes).join("circle").attr("r", 10).attr("class", "node");

        simulation.on("tick", () => {
            link.attr("x1", d => d.source.x).attr("y1", d => d.source.y)
                .attr("x2", d => d.target.x).attr("y2", d => d.target.y);
            node.attr("cx", d => d.x).attr("cy", d => d.y);
        });
    </script>
</body>
</html>
`

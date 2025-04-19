let map = L.map("map").setView([20, 0], 2);
L.tileLayer("https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png").addTo(map);
let geojsonLayer = null;

function updateMap(geojsonText) {
  const geojson = JSON.parse(geojsonText);

  if (geojsonLayer) {
    map.removeLayer(geojsonLayer);
  }

  geojsonLayer = L.geoJSON(geojson).addTo(map);
  map.fitBounds(geojsonLayer.getBounds());
}

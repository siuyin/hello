var wwd = new WorldWind.WorldWindow("canvasOne");

wwd.addLayer(new WorldWind.BMNGOneImageLayer());
wwd.addLayer(new WorldWind.BMNGLandsatLayer());

//wwd.addLayer(new WorldWind.CompassLayer());
wwd.addLayer(new WorldWind.CoordinatesDisplayLayer(wwd));
//wwd.addLayer(new WorldWind.ViewControlsLayer(wwd));


// place marks
var placemarkLayer = new WorldWind.RenderableLayer("Placemark");
wwd.addLayer(placemarkLayer);

var placemarkAttributes = new WorldWind.PlacemarkAttributes(null);
placemarkAttributes.imageOffset = new WorldWind.Offset(
    WorldWind.OFFSET_FRACTION, 0.3,
        WorldWind.OFFSET_FRACTION, 0.0);
placemarkAttributes.labelAttributes.color = WorldWind.Color.YELLOW;
placemarkAttributes.labelAttributes.offset = new WorldWind.Offset(
    WorldWind.OFFSET_FRACTION, 0.5,
        WorldWind.OFFSET_FRACTION, 1.0);
placemarkAttributes.imageSource = WorldWind.configuration.baseUrl + "images/pushpins/plain-red.png";

// position place mark
var position = new WorldWind.Position(55.0, -106.0, 100.0);
var placemark = new WorldWind.Placemark(position, false, placemarkAttributes);
placemark.label = "Placemark\n" +
    "Lat " + placemark.position.latitude.toPrecision(4).toString() + "\n" +
        "Lon " + placemark.position.longitude.toPrecision(5).toString();
        placemark.alwaysOnTop = true;
placemarkLayer.addRenderable(placemark);


// adding polygons
var polygonLayer = new WorldWind.RenderableLayer();
wwd.addLayer(polygonLayer);

var polygonAttributes = new WorldWind.ShapeAttributes(null);
polygonAttributes.interiorColor = new WorldWind.Color(0, 1, 1, 0.75);
polygonAttributes.outlineColor = WorldWind.Color.BLUE;
polygonAttributes.drawOutline = true;
polygonAttributes.applyLighting = true;
var boundaries = [];
boundaries.push(new WorldWind.Position(20.0, -75.0, 7000.0));
boundaries.push(new WorldWind.Position(25.0, -85.0, 700000.0));
boundaries.push(new WorldWind.Position(20.0, -95.0, 700000.0));
var polygon = new WorldWind.Polygon(boundaries, polygonAttributes);
polygon.extrude = true;
polygonLayer.addRenderable(polygon);

// Add WMS imagery
var serviceAddress = "https://neo.sci.gsfc.nasa.gov/wms/wms?SERVICE=WMS&REQUEST=GetCapabilities&VERSION=1.3.0";
//var layerName = "MOD_LSTD_CLIM_M";
var layerName = "MOP_CO_M";

var createLayer = function (xmlDom) {
    var wms = new WorldWind.WmsCapabilities(xmlDom);
    var wmsLayerCapabilities = wms.getNamedLayer(layerName);
    var wmsConfig = WorldWind.WmsLayer.formLayerConfiguration(wmsLayerCapabilities);
    var wmsLayer = new WorldWind.WmsLayer(wmsConfig);
    wwd.addLayer(wmsLayer);
};

var logError = function (jqXhr, text, exception) {
    console.log("There was a failure retrieving the capabilities document: " +
        text +
    " exception: " + exception);
};

//$.get(serviceAddress).done(createLayer).fail(logError);

// Keep track of the DOM element we'll use to show what's selected.
var pickResult = document.getElementById("pick-result");

// The common pick-handling function.
var handlePick = function(o) {
  // The input argument is either an Event or a TapRecognizer. Both have the same properties for determining
  // the mouse or tap location.
  var x = o.clientX,
      y = o.clientY;

  // Perform the pick. Must first convert from window coordinates to canvas coordinates, which are
  // relative to the upper left corner of the canvas rather than the upper left corner of the page.
  var pickList = wwd.pick(wwd.canvasCoordinates(x, y));

  // Report the top picked object, if any.
  var topPickedObject = pickList.topPickedObject();
  if (topPickedObject && topPickedObject.isTerrain) {
    pickResult.textContent = topPickedObject.position;
  } else if (topPickedObject) {
    pickResult.textContent = topPickedObject.userObject.displayName;
  } else {
    pickResult.textContent = "Nothing Selected";
  }
};

// Listen for mouse moves and touch taps.
//wwd.addEventListener("mousemove", handlePick);
wwd.addEventListener("click", handlePick);

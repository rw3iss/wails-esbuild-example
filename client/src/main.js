require('core-js/stable');
const runtime = require('@wailsapp/runtime');

// Main entry point
function start() {

	var mystore = runtime.Store.New('Counter');

	// Ensure the default app div is 100% wide/high
	var app = document.getElementById('app');
	app.style.width = '100%';
	app.style.height = '100%';

	// Inject html
	app.innerHTML = `
        <div class='logo'></div>
        <div class='container'>
        The app!
        </div>
	`;
};

// We provide our entrypoint as a callback for runtime.Init
runtime.Init(start);
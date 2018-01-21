import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import App from './app/App';
import registerServiceWorker from './registerServiceWorker';

const RECIPE = {
	units: 'g',
	lye_type: 'naoh',
	lipid_weight: 40,
	water_lipid_ratio: 0.25,
	super_fat_percentage: 0.04,
	fragrance_ratio: 0.04,
	lipids: [
		{name: 'Olive Oil', sap: '.123', weight: '10', percentage: '50'},
 		{name: 'Coconut Oil', sap: '.145', weight: '5', percentage: '25'},
		{name: 'Palm Oil', sap: '.167', weight: '5', percentage: '25'},
	]
}

ReactDOM.render(<App />, document.getElementById('root'));
registerServiceWorker();

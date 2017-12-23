import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import SoapCalc from './App';
import registerServiceWorker from './registerServiceWorker';


const RECIPE = {
	units: 'oz',
	lye_type: 'naoh',
	lipid_weight: 40,
	water_to_lipid_ratio: 0.35,
	super_fat_percentage: 0.05,
	fragrance_ratio: 0.05,
	lipids: [
		{name: 'Olive Oil', sap: '.123', weight: '10', percentage: '50'},
 		{name: 'Coconut Oil', sap: '.145', weight: '5', percentage: '25'},
		{name: 'Palm Oil', sap: '.167', weight: '5', percentage: '25'},
	]
}

ReactDOM.render(<SoapCalc recipe={RECIPE}/>, document.getElementById('root'));
registerServiceWorker();

import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import SoapCalc from './App';
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


const LIPIDS = [
	{"name": "Olive Oil", "description": "Olive Oil", "scientific_name": "", "naoh": 0.135, "koh": 0.19, "iodine": 85, "ins": 105, "lauric": 0, "myristic": 0, "palmitic": 0.14, "stearic": 0.03, "ricinoleic": 0, "oleic": 0.69, "linoleic": 0.12, "linolenic": 0.01, "hardness": 17, "cleansing": 0, "condition": 82, "bubbly": 0, "creamy": 17},
	{"name": "Coconut Oil, 76 deg", "description": "Coconut Oil, 76 deg", "scientific_name": "", "naoh": 0.183, "koh": 0.257, "iodine": 10, "ins": 258, "lauric": 0.48, "myristic": 0.19, "palmitic": 0.09, "stearic": 0.03, "ricinoleic": 0, "oleic": 0.08, "linoleic": 0.02, "linolenic": 0, "hardness": 0, "cleansing": 0, "condition": 0, "bubbly": 0, "creamy": 0},
	{"name": "Palm Oil", "description": "Palm Oil", "scientific_name": "", "naoh": 0.142, "koh": 0.199, "iodine": 53, "ins": 145, "lauric": 0, "myristic": 0.01, "palmitic": 0.44, "stearic": 0.05, "ricinoleic": 0, "oleic": 0.39, "linoleic": 0.1, "linolenic": 0, "hardness": 50, "cleansing": 1, "condition": 49, "bubbly": 1, "creamy": 49},
	{"name": "Castor Oil", "description": "Castor Oil", "scientific_name": "", "naoh": 0.128, "koh": 0.18, "iodine": 86, "ins": 95, "lauric": 0, "myristic": 0, "palmitic": 0, "stearic": 0, "ricinoleic": 0.9, "oleic": 0.04, "linoleic": 0.04, "linolenic": 0, "hardness": 0, "cleansing": 0, "condition": 98, "bubbly": 90, "creamy": 90},
	{"name": "Hemp Oil", "description": "Hemp Oil", "scientific_name": "", "naoh": 0.138, "koh": 0.193, "iodine": 165, "ins": 39, "lauric": 0, "myristic": 0, "palmitic": 0.06, "stearic": 0.02, "ricinoleic": 0, "oleic": 0.12, "linoleic": 0.57, "linolenic": 0.21, "hardness": 8, "cleansing": 0, "condition": 90, "bubbly": 0, "creamy": 8}
]

ReactDOM.render(<SoapCalc lipids={LIPIDS}/>, document.getElementById('root'));
registerServiceWorker();

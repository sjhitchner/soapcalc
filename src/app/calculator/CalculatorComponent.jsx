import React, { Component } from 'react';
import Autosuggest from 'react-autosuggest';
import { 
  Button,
	Col,
	Grid,
	Row,
} from 'react-bootstrap';
/*
import {
  FragranceRatioInput,
  LyeType,
  LipidWeight,
  PercentageInput,
  SuperFatInput,
  Units,
  UnitsInput,
  WaterLipidRatioInput,
} from '../common/index';
*/
import FragranceRatioInput from '../common/FragranceRatioInputComponent';
import LyeType from '../common/LyeTypeComponent';
import LipidWeight from '../common/LipidWeightComponent';
import PercentageInput from '../common/PercentageInputComponent';
import SuperFatInput from '../common/SuperFatInputComponent';
import Units from '../common/UnitsComponent';
import UnitsInput from '../common/UnitsInputComponent';
import WaterLipidRatioInput from '../common/WaterLipidRatioInputComponent';

const LIPIDS = [
	{name: "Olive Oil", "description": "Olive Oil", "scientific_name": "", naoh: 0.135, "koh": 0.19, "iodine": 85, "ins": 105, "lauric": 0, "myristic": 0, "palmitic": 0.14, "stearic": 0.03, "ricinoleic": 0, "oleic": 0.69, "linoleic": 0.12, "linolenic": 0.01, "hardness": 17, "cleansing": 0, "condition": 82, "bubbly": 0, "creamy": 17},
	{name: "Coconut Oil, 76 deg", "description": "Coconut Oil, 76 deg", "scientific_name": "", naoh: 0.183, "koh": 0.257, "iodine": 10, "ins": 258, "lauric": 0.48, "myristic": 0.19, "palmitic": 0.09, "stearic": 0.03, "ricinoleic": 0, "oleic": 0.08, "linoleic": 0.02, "linolenic": 0, "hardness": 0, "cleansing": 0, "condition": 0, "bubbly": 0, "creamy": 0},
	{name: "Palm Oil", "description": "Palm Oil", "scientific_name": "", naoh: 0.142, "koh": 0.199, "iodine": 53, "ins": 145, "lauric": 0, "myristic": 0.01, "palmitic": 0.44, "stearic": 0.05, "ricinoleic": 0, "oleic": 0.39, "linoleic": 0.1, "linolenic": 0, "hardness": 50, "cleansing": 1, "condition": 49, "bubbly": 1, "creamy": 49},
	{name: "Castor Oil", "description": "Castor Oil", "scientific_name": "", naoh: 0.128, "koh": 0.18, "iodine": 86, "ins": 95, "lauric": 0, "myristic": 0, "palmitic": 0, "stearic": 0, "ricinoleic": 0.9, "oleic": 0.04, "linoleic": 0.04, "linolenic": 0, "hardness": 0, "cleansing": 0, "condition": 98, "bubbly": 90, "creamy": 90},
	{name: "Hemp Oil", "description": "Hemp Oil", "scientific_name": "", naoh: 0.138, "koh": 0.193, "iodine": 165, "ins": 39, "lauric": 0, "myristic": 0, "palmitic": 0.06, "stearic": 0.02, "ricinoleic": 0, "oleic": 0.12, "linoleic": 0.57, "linolenic": 0.21, "hardness": 8, "cleansing": 0, "condition": 90, "bubbly": 0, "creamy": 8}
]

//lipids={LIPIDS} 

function sortLipids(lipids) {
  return lipids; //lipids.sort((a,b) => a.percentage < b.percentage);
}

class Calculator extends Component {
	constructor(props) {
		super(props);
		this.state = {
			lipidSearch: '',
			units: 'oz',
			lyeType: 'naoh',
			lipidWeight: 40,
			waterLipidRatio: 0.35,
			superFatPercentage: 0.05,
			fragranceRatio: 0.05,
			selectedLipids:[],
		};
	}

	handleUnitsChange = (newUnits) => {
		this.setState({
			units: newUnits,
		});
	}
  
	handleLyeTypeChange = (newLyeType) => {
		this.setState({
			lyeType: newLyeType,
		});
	}

	handleLipidWeightChange = (newLipidWeight) => {
		this.setState({
			lipidWeight: newLipidWeight,
		});
	}

	handleWaterLipidRatioChange = (newWaterLipidRatio) => {
		this.setState({
			waterLipidRatio: newWaterLipidRatio,
		});
	}

	handleSuperFatPercentageChange = (newSuperFat) => {
		this.setState({
			superFatPercentage: newSuperFat,
		});
	}

	handleFragranceRatioChange = (newFragranceRatio) => {
		this.setState({
			fragranceRatio: newFragranceRatio,
		});
	}

	handleAddLipid = (lipid) => {
    for (let l of this.state.selectedLipids) {
      if (l.name === lipid.name) {
         return;
      }
    }

		lipid.weight = 0;
		lipid.percentage = 0;

		this.setState({
			selectedLipids: sortLipids(this.state.selectedLipids.concat([lipid])),
		});
	};

	handleDeleteLipid = (index) => {
		this.setState({
			selectedLipids: sortLipids(this.state.selectedLipids.filter((s, idx) => idx !== index)),
		});
	}

	handleLipidUpdate = (idx, type, value) => {
		let lipids = this.state.selectedLipids.slice()

		switch (type) {
			case 'weight':
        lipids[idx].weight = value;
        lipids[idx].percentage = value / this.state.lipidWeight;	
				break;
			case 'percentage':
        lipids[idx].weight = this.state.lipidWeight * value;
        lipids[idx].percentage = value;
				break;

			default:
				break;
		}

		this.setState({
			selectedLipids: sortLipids(lipids)
		});
	}

	handleSubmit = (e) => {
		const recipe = {
			units: this.state.units,
			lye_type: this.state.lyeType,
			lipid_weight: this.state.lipidWeight,
			water_lipid_ratio: this.state.waterLipidRatio,
			super_fat_percentage: this.state.superFatPercentage,
			fragrance_ratio: this.state.fragranceRatio, 
			lipids: this.state.selectedLipids,
		}
	
		alert(`Recipe: ` + JSON.stringify(recipe));
	}

	render() {
		return (
			<Grid>
				<Row>
					<CalculatorParameters 
						units={this.state.units}
						onUnitsChange={this.handleUnitsChange}
						lyeType={this.state.lyeType}
						onLyeTypeChange={this.handleLyeTypeChange}
						lipidWeight={this.state.lipidWeight}
						onLipidWeightChange={this.handleLipidWeightChange}
						waterLipidRatio={this.state.waterLipidRatio}
						onWaterLipidRatioChange={this.handleWaterLipidRatioChange}
						superFatPercentage={this.state.superFatPercentage}
						onSuperFatPercentageChange={this.handleSuperFatPercentageChange}
						fragranceRatio={this.state.fragranceRatio}
						onFragranceRatioChange={this.handleFragranceRatioChange}
					/>
				</Row>
				<Row>
					<LipidSelection
						allLipids={LIPIDS} //this.props.lipids}
						selectedLipids={this.state.selectedLipids}
						addLipid={this.handleAddLipid}
						updateLipid={this.handleLipidUpdate}
						deleteLipid={this.handleDeleteLipid}
						totalWeight={this.state.lipidWeight}
						units={this.state.units}
					/>
				</Row>
				<Row>
					<Button bsStyle="primary" bsSize="large" disabled>Submit Recipe</Button>
				</Row>
			</Grid>
		);
	}
}

class CalculatorParameters extends Component {
	render() {
		return (
			<Grid>
				<Row className="show-grid">
					<Col xs={6} md={6}>
						<LyeType 
							value={this.props.lyeType} 
							onChange={this.props.onLyeTypeChange} />
					</Col>
					<Col xs={6} md={6}>
						<Units 
							value={this.props.units}
							onChange={this.props.onUnitsChange} />
					</Col>
				</Row>
				<Row>
					<Col xs={6} md={6}>
						<LipidWeight 
							value={this.props.lipidWeight}
							units={this.props.units} 
							onChange={this.props.onLipidWeightChange} />
					</Col>
					<Col xs={6} md={6}>
						<SuperFatInput
							value={this.props.superFatPercentage}
							onChange={this.props.onSuperFatPercentageChange} />
					</Col>
				</Row>
				<Row>
					<Col xs={6} md={6}>
						<WaterLipidRatioInput
							value={this.props.waterLipidRatio}
							onChange={this.props.onWaterLipidRatioChange} />
					</Col>
					<Col xs={6} md={6}>
						<FragranceRatioInput
							value={this.props.fragranceRatio}
							onChange={this.props.onFragranceRatioChange} />
					</Col>
				</Row>
			</Grid>
		);
	}	
}

class LipidSelection extends Component {
	render() {
		return (
			<Grid>
				<Row>
					<LipidLookup
						allLipids={this.props.allLipids}
						onSelected={this.props.addLipid}
					/>
				</Row>
				<Row>
					<LipidTable
						selectedLipids={this.props.selectedLipids}
						updateLipid={this.props.updateLipid}
						deleteLipid={this.props.deleteLipid}
						totalWeight={this.props.totalWeight}
						units={this.props.units}
					/>
    			</Row>
			</Grid>
		);
	}	
}

// Use your imagination to render suggestions.
const renderSuggestion = suggestion => (
	<div>
		{suggestion.name}
	</div>
);

class LipidLookup extends Component {
	constructor(props) {
		super(props);

		this.state = {
			value: '',
			suggestions: [],
		};
	}

	onSuggestionSelected = (event, { suggestion }) => {
		this.props.onSelected(suggestion);
		this.setState({
			value: '',
		});
	};

	onChange = (event, { newValue }) => {
		this.setState({
			value: newValue
		});
	};

	// Autosuggest will call this function every time you need to update suggestions.
	// You already implemented this logic above, so just use it.
	onSuggestionsFetchRequested = ({ value }) => {
		this.setState({
			suggestions: this.getSuggestions(value)
		});
	};

	// Autosuggest will call this function every time you need to clear suggestions.
	onSuggestionsClearRequested = () => {
		this.setState({
			suggestions: []
		});
	};

	getSuggestionValue = suggestion => {
		return suggestion.name;
	}

	getSuggestions = value => {
		const inputValue = value.trim().toLowerCase();
		const inputLength = inputValue.length;

		return inputLength === 0 ? [] : this.props.allLipids.filter(lipid =>
			lipid.name.toLowerCase().slice(0, inputLength) === inputValue
		);
	};

	render() {
		const { value, suggestions } = this.state;

		// Autosuggest will pass through all these props to the input.
		const inputProps = {
			placeholder: 'Select a Lipid/Oil',
			value,
			onChange: this.onChange
		};

		return (
			<Autosuggest
				suggestions={suggestions}
				onSuggestionSelected={this.onSuggestionSelected}
				onSuggestionsFetchRequested={this.onSuggestionsFetchRequested}
				onSuggestionsClearRequested={this.onSuggestionsClearRequested}
				getSuggestionValue={this.getSuggestionValue}
				renderSuggestion={renderSuggestion}
				inputProps={inputProps}
			/>
		);
	}
}

class LipidTable extends Component {

	handleWeightChange = (idx) => (value) => {
		this.props.updateLipid(idx, 'weight', value);
	}

	handlePercentageChange = (idx) => (value) => {
		this.props.updateLipid(idx, 'percentage', value);
	}

	handleDelete = (idx) => (e) => {
		this.props.deleteLipid(idx);
	}

	render() {
		let sumWeight = 0;
		let sumPercentage = 0;

		return (
			<table className="table">
				<thead>
    				<tr>
      					<th scope="col">#</th>
      					<th scope="col">Lipid</th>
      					<th scope="col">SAP</th>
      					<th scope="col">%</th>
      					<th scope="col">Weight</th>
      					<th scope="col"></th>
					</tr>
				</thead>
				<tbody>
          {this.props.selectedLipids.map((lipid, idx) => {
 							sumWeight += lipid.weight;
 							sumPercentage += lipid.percentage;
 
 							return (
 								<LipidTableRow
 									key={idx}
 				        	num={idx+1}
 									name={lipid.name}
 									sap={lipid.naoh}
 									percentage={lipid.percentage}
 									weight={lipid.weight}
 									units={this.props.units}
 									onWeightUpdate={this.handleWeightChange(idx)}
 									onPercentageUpdate={this.handlePercentageChange(idx)}
 									onDelete={this.handleDelete(idx)}
 								/>
 							)
 						})}
						<tr>
							<th colSpan="3">Total Weight</th>
							<th>{sumPercentage*100}</th>
							<th>{sumWeight*1}&nbsp;{this.props.units}</th>
              <th>&nbsp;</th>
						</tr>
				</tbody>
			</table>
		);
	}
}

class LipidTableRow extends Component {

	render() {
		return (
			<tr> 
				<th scope="row">{this.props.num}</th>
      			<td>{this.props.name}</td>
      			<td>{this.props.sap}</td>
      			<td>
					<PercentageInput
						name="lipid_percentage"
						value={this.props.percentage}
						onChange={this.props.onPercentageUpdate}
					/>
				</td>
      			<td>
					<UnitsInput 
						name="lipid_weight"
						value={this.props.weight}
						units={this.props.units}
						placeholder={this.props.placeholder}
						onChange={this.props.onWeightUpdate}
					/>
				</td>
				<td>
					<Button
						bsSize="small"
						onClick={this.props.onDelete} >
					delete
					</Button>
				</td>
    		</tr>
		);
	}
}

export default Calculator;

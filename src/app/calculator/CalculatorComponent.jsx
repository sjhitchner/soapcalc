import React, { Component } from 'react';
import Autosuggest from 'react-autosuggest';
import { 
	Col,
	Grid,
	Row,
} from 'react-bootstrap';
import {
  Units,
  LyeType,
  LipidWeight,
  WaterLipidRatioInput,
  SuperFatInput,
  FragranceRatioInput,
} from '../common/CommonComponent';

function sortLipids(lipids) {
   const hash = {};
   return lipids
     .sort((a,b) => a.percentage < b.percentage)
     .filter(function(lipid) {
         var name = lipid.name;
         return !hash[name] && (hash[name] = true);
   });
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

	handleLipidUpdate = (index, type, value) => {
		var lipids = this.state.selectedLipids.slice()
		//lipids[index].weight = weightpercentage;
		//lipids[index].percentage = percentage;
		switch (type) {
			case 'weight':
				
				break;
			case 'percentage':
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
						allLipids={this.props.lipids}
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
	render() {
		const rows = [];
		var sumWeight = 0;
		var sumPercentage = 0;

		this.props.selectedLipids.forEach((lipid, i) => {
			sumWeight += lipid.weight;
			sumPercentage += lipid.percentage;
			rows.push(
				<LipidTableRow
				        key={lipid.name}
				        num={i}
						name={lipid.name}
						sap={lipid.naoh}
						percentage={lipid.percentage}
						weight={lipid.weight}
						units={this.props.units}
						onUpdate={this.props.updateLipid}
						onDelete={this.props.deleteLipid}
				/>
			);
		});

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
						{rows}
						<tr>
							<th colSpan="3">Total Weight</th>
							<th>{sumPercentage}</th>
							<th>{sumWeight}&nbsp;{this.props.units}</th>
						</tr>
				</tbody>
			</table>
		);
	}
}

class LipidTableRow extends Component {

	handleWeightChange = (value) => {
		this.props.onUpdate(this.props.num, 'weight', value);
	}

	handlePercentageChange = (value) => {
		this.props.onUpdate(this.props.num, 'percentage', value);
	}

	handleDelete = (e) => {
		this.props.deleteLipid(this.props.num);
	}

	render() {
		return (
			<tr key={this.props.key}>
				<th scope="row">{this.props.num+1}</th>
      			<td>{this.props.name}</td>
      			<td>{this.props.sap}</td>
      			<td>
					<PercentageInput
						name="lipid_percentage"
						value={this.props.percentage}
						onChange={this.handleWeightChange}
					/>
				</td>
      			<td>
					<UnitsInput 
						name="lipid_weight"
						value={this.props.value}
						units={this.props.units}
						placeholder={this.props.placeholder}
						onChange={this.handlePercentageChange}
					/>
				</td>
				<td>
					<Button
						bsSize="small"
						onClick={this.handleDelete} >
					delete
					</Button>
				</td>
    		</tr>
		);
	}
}

export default Calculator;

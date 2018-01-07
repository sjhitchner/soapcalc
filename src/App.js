import React, { Component } from 'react';
import Autosuggest from 'react-autosuggest';
import { 
	Button,
	ButtonToolbar,
	ControlLabel,
	FormControl,
	FormGroup,
	Grid,
	HelpBlock,
	InputGroup,
	Row,
	Col,
	ToggleButton, 
	ToggleButtonGroup, 
} from 'react-bootstrap';
//import logo from './logo.svg';
import './App.css';
import 'bootstrap/dist/css/bootstrap.min.css';

/*
<!-- Latest compiled and minified CSS -->
<link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/latest/css/bootstrap.min.css">

<!-- Optional theme -->
<link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/latest/css/bootstrap-theme.min.css">

      <div className="SoapCalc">
        <header className="SoapCalc-header">
          <img src={logo} className="SoapCalc-logo" alt="logo" />
          <h1 className="SoapCalc-title">Welcome to React</h1>
        </header>
        <p className="SoapCalc-intro">
          To get started, edit <code>src/SoapCalc.js</code> and save to reload.
        </p>
		<LyeSelector />
		<br />
		<UnitsSelector />
		<br />
		<LipidWeight />
		<br />
		<LipidList />
      </div>
    	);
*/

function sortLipids(lipids) {
   return lipids.sort((a,b) => a.percentage < b.percentage);
}

class SoapCalc extends Component {
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
        for (var l of this.state.selectedLipids) {
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
		var lipids = this.state.selectedLipids.slice()

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
					<SoapCalcParameters 
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
					<SoapCalcLipidSelection
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

class SoapCalcParameters extends Component {
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

class SoapCalcLipidSelection extends Component {
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
		var sumWeight = 0;
		var sumPercentage = 0;

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

class Units extends Component {
  
	handleChange = (e) => {
		this.props.onChange(e);
	}

	render() {
		return (
			<FormGroup controlId="units">
				<ControlLabel>Units</ControlLabel>
				<ButtonToolbar>
					<ToggleButtonGroup
						name="units"
						type="radio"
						onChange={this.handleChange}
						defaultValue={this.props.value}>
						<ToggleButton value={'oz'}>Ounces</ToggleButton>
						<ToggleButton value={'g'}>Grams</ToggleButton>
					</ToggleButtonGroup>
				</ButtonToolbar>
				<FormControl.Feedback />
			</FormGroup>
		)
	}
}

class LyeType extends Component {
  
	handleChange = (e) => {
		this.props.onChange(e);
	}

	render() {
		return (
			<FormGroup controlId="lyeType">
   				<ControlLabel>Lye Type</ControlLabel>
				<ButtonToolbar>
					<ToggleButtonGroup
						name="lye_type"
						type="radio"
						onChange={this.handleChange}
						defaultValue={this.props.value}>
						<ToggleButton value={'naoh'}>NaOH</ToggleButton>
						<ToggleButton value={'koh'} disabled>KOH</ToggleButton>
					</ToggleButtonGroup>
				</ButtonToolbar>
				<FormControl.Feedback />
			</FormGroup>
		)
	}
}

class LipidWeight extends Component {
  
	handleChange = (e) => {
		this.props.onChange(e.target.value);
	}

	getValidationState () {
		/*
    	const length = this.state.value.length;
    	if (length > 10) return 'success';
    	else if (length > 5) return 'warning';
    	else if (length > 0) return 'error';
		*/
	}

	render() {
		return (
			<FormGroup
				controlId="lipidWeight"
				validationState={this.getValidationState()}>
      			<ControlLabel>Lipid Weight</ControlLabel>
				<UnitsInput 
					name="lipid_weight"
					value={this.props.value}
					units={this.props.units}
					placeholder={this.props.placeholder} />
				<FormControl.Feedback />
            	<HelpBlock>Validation is based on string length.</HelpBlock>
			</FormGroup>
		)
	}
}

class WaterLipidRatioInput extends Component {

	render() {
		return (
			<FormGroup controlId="waterLipidRatio">
   				<ControlLabel>Water to Lipid Ratio</ControlLabel>
				<PercentageInput
					name="water_lipid_ratio"
					value={this.props.value} 
					onChange={this.props.onChange} />
				<FormControl.Feedback />
            	<HelpBlock>Validation is based on string length.</HelpBlock>
			</FormGroup>
		)
	}
}

class SuperFatInput extends Component {

	render() {
		return (
			<FormGroup controlId="superFatPercentage">
   				<ControlLabel>Super Fat Percentage</ControlLabel>
				<PercentageInput 
					name="super_fat_percentage"
					value={this.props.value}
					onChange={this.props.onChange} />
				<FormControl.Feedback />
            	<HelpBlock>Validation is based on string length.</HelpBlock>
			</FormGroup>
		)
	}
}

class FragranceRatioInput extends Component {

	render() {
		return (
			<FormGroup controlId="fragranceRatio">
   				<ControlLabel>Fragrance Ratio</ControlLabel>
				<PercentageInput
					name="fragrance_ratio"
					value={this.props.value}
					onChange={this.props.onChange} />
				<FormControl.Feedback />
            	<HelpBlock>Validation is based on string length.</HelpBlock>
			</FormGroup>
		)
	}
}

class UnitsInput extends Component {

	handleChange = (e) => {
		this.props.onChange(e.target.value);
	}

	render() {
		return (
			<InputGroup>
      			<FormControl
					type="text"
					bsSize="sm"
					placeholder={this.props.placeholder}
			 		onChange={this.handleChange}
					value={this.props.value} />
				<InputGroup.Addon>{this.props.units}</InputGroup.Addon>
			</InputGroup>
		)
	}
}

class PercentageInput extends Component {
  
	handleChange = (e) => {
    const value = e.target.value;
    if (value >= 0 && value <= 100) {
		    this.props.onChange(value/100);
    }
	}

	/*
	getValidationState () {
		const value = this.state.value;
		if (value > 0 && value <= 1 {
			return 'success';
		}
    	else if (value > 5) {
			return 'warning';
		}
    	else {
			return 'error';
		}
	}
	*/

	render() {
		return (
			<InputGroup>
      			<FormControl
					type="text"
					bsSize="sm"
					placeholder={this.props.placeholder}
			 		onChange={this.handleChange}
					value={this.props.value * 100} />
				<InputGroup.Addon>%</InputGroup.Addon>
			</InputGroup>
		)
	}
}


/*
class LipidList extends Component {
	constructor(props) {
		super(props);
	}

	render() {
		var lipids = [
			'<div>Olive Oil</div>',
		];

		return (
			<div>
				<LipidSelector />
				<div className="lipids">{lipids}</div>
			</div>
		);
	}
}

class LipidSelector extends Component {
	constructor(props) {
		super(props);
		this.state = {
			value: '',
		};
	}

	getValidationState() {
		const length = this.state.value.length;
		if (length > 10) return 'success';
		else if (length > 5) return 'warning';
		else if (length > 0) return 'error';
		return null;
	}

	handleChange(e) {
		this.setState({ value: e.target.value });
	}

	render() {
		return (
			<FormGroup controlId="formBasicText" validationState={this.getValidationState()}>
				<ControlLabel>Search for Lipid</ControlLabel>
				<FormControl type="text" value={this.state.value} placeholder="Enter text" onChange={this.handleChange} />
				<FormControl.Feedback />
				<HelpBlock>Validation is based on string length.</HelpBlock>
			</FormGroup>
		);
	}
}
*/
/*
var BoardSwitcher = React.createClass({
  render: function() {
    var boards = [];
    for (var ii = 0; ii < this.props.numBoards; ii++) {
      var isSelected = ii === 0;
      boards.push(
        <Board index={ii} selected={isSelected} key={ii} />
      );
    }
    
    return (
      <div>
        <div className="boards">{boards}</div>
        <button>Toggle</button>
      </div>
    );
  }
});
*/


export default SoapCalc;

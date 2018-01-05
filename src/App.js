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

		this.handleUnitsChange = this.handleUnitsChange.bind(this);
    	this.handleLyeTypeChange = this.handleLyeTypeChange.bind(this);
		this.handleLipidWeightChange = this.handleLipidWeightChange.bind(this);
		this.handleWaterLipidRatioChange = this.handleWaterLipidRatioChange.bind(this);
		this.handleSuperFatPercentageChange = this.handleSuperFatPercentageChange.bind(this);
		this.handleFragranceRatioChange = this.handleFragranceRatioChange.bind(this);
		this.handleAddLipid = this.handleAddLipid.bind(this);
		this.handleDeleteLipid = this.handleDeleteLipid.bind(this);
		this.handleLipidPercentageUpdate = this.handleLipidPercentageUpdate.bind(this);  
	}

	handleUnitsChange(newUnits) {
		this.setState({
			units: newUnits,
		});
	}
  
	handleLyeTypeChange(newLyeType) {
		this.setState({
			lyeType: newLyeType,
		});
	}

	handleLipidWeightChange(newLipidWeight) {
		this.setState({
			lipidWeight: newLipidWeight,
		});
	}

	handleWaterLipidRatioChange(newWaterLipidRatio) {
		this.setState({
			waterLipidRatio: newWaterLipidRatio,
		});
	}

	handleSuperFatPercentageChange(newSuperFat) {
		this.setState({
			superFatPercentage: newSuperFat,
		});
	}

	handleFragranceRatioChange(newFragranceRatio) {
		this.setState({
			fragranceRatio: newFragranceRatio,
		});
	}

	handleAddLipid(lipid) {
		lipid.percentage = 0;

		var lipids = this.state.selectedLipids.slice()
		lipids.push(lipid)

		this.setState({
			selectedLipids: lipids
		});
	};

	handleDeleteLipid(index) {
		var lipids = this.state.selectedLipids.slice()
		lipids.splice(index, 1);

		this.setState({
			selectedLipids: lipids
		});
	}

	handleLipidPercentageUpdate(index, percentage) {
		var lipids = this.state.selectedLipids.slice()
		lipids[index].percentage = percentage;

		this.setState({
			selectedLipids: lipids
		});
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
						updatePercentage={this.handleLipidPercentageUpdate}
						deleteLipid={this.handleDeleteLipid}
						totalWeight={this.state.lipidWeight}
					/>
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
						updatePercentage={this.props.updatePercentage}
						deleteLipid={this.props.deleteLipid}
						totalWeight={this.props.totalWeight}
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

		this.props.selectedLipids.forEach((lipid, i) => {
			const weight = this.props.totalWeight * lipid.percentage;
			sumWeight += weight;
			rows.push(
				<LipidTableRow
				        key={i}
				        num={i}
						name={lipid.name}
						sap={lipid.naoh}
						percentage={lipid.percentage}
						updatePercentage={this.props.updatePercentage}
						weight={weight}
						deleteLipid={this.props.deleteLipid}
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
							<th colSpan="4">Total Weight</th>
							<th>{sumWeight}</th>
						</tr>
				</tbody>
			</table>
		);
	}
}

class LipidTableRow extends Component {
	constructor(props) {
		super(props);
		this.handleChange = this.handleChange.bind(this);
		this.handleDelete = this.handleDelete.bind(this);
	}

	handleChange(value) {
		this.props.updatePercentage(this.props.num, value);
	}

	handleDelete(e) {
		this.props.deleteLipid(this.props.num);
	}

	render() {
		return (
			<tr>
				<th scope="row">{this.props.num+1}</th>
      			<td>{this.props.name}</td>
      			<td>{this.props.sap}</td>
      			<td>
					<PercentageInput
						name="lipid_percentage"
						value={this.props.percentage}
						onChange={this.handleChange}
					/>
				</td>
      			<td>{this.props.weight}</td>
				<td>
					<Button
						bsSize="small"
						bsStyle="primary"
						onClick={this.handleDelete} >
					delete
					</Button>
				</td>
    		</tr>
		);
	}
}

class Units extends Component {
	constructor(props) {
		super(props);
		this.handleChange = this.handleChange.bind(this);
	}
  
	handleChange(e) {
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
	constructor(props) {
		super(props);
		this.handleChange = this.handleChange.bind(this);
	}
  
	handleChange(e) {
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
						<ToggleButton value={'koh'}>KOH</ToggleButton>
					</ToggleButtonGroup>
				</ButtonToolbar>
				<FormControl.Feedback />
			</FormGroup>
		)
	}
}

class LipidWeight extends Component {
	constructor(props) {
		super(props);
		this.handleChange = this.handleChange.bind(this);
	}
  
	handleChange(e) {
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
				<InputGroup>
					<FormControl
						type="text"
						bsSize="sm"
						placeholder={this.props.placeholder}
			 			onChange={this.handleChange}
						defaultValue={this.props.value} />
					<InputGroup.Addon>{this.props.units}</InputGroup.Addon>
				</InputGroup>
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

class PercentageInput extends Component {
	constructor(props) {
		super(props);
		this.handleChange = this.handleChange.bind(this);
	}
  
	handleChange(e) {
		this.props.onChange(e.target.value/100);
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
					defaultValue={this.props.value * 100} />
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

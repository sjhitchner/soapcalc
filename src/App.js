import React, { Component } from 'react';
import { 
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
	render() {
		return (
			<Grid>
				<Row>
					<SoapCalcParameters 
						units={this.props.recipe.units}
						lye_type={this.props.recipe.lye_type}
						lipid_weight={this.props.recipe.lipid_weight}
						water_lipid_ratio={this.props.recipe.water_lipid_ratio}
						super_fat_percentage={this.props.recipe.super_fat_percentage}
						fragrance_ratio={this.props.recipe.fragrance_ratio}
					/>
				</Row>
				<Row>
					<SoapCalcLipidSelection lipids={this.props.recipe.lipids}/>
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
						<LyeType value={this.props.lye_type} />
					</Col>
					<Col xs={6} md={6}>
						<Units value={this.props.units} />
					</Col>
				</Row>
				<Row>
					<Col xs={6} md={6}>
						<LipidWeight value={this.props.lipid_weight} units={this.props.units} />
					</Col>
					<Col xs={6} md={6}>
						<SuperFatInput value={this.props.super_fat_percentage} />
					</Col>
				</Row>
				<Row>
					<Col xs={6} md={6}>
						<WaterLipidRatioInput value={this.props.water_lipid_ratio} />
					</Col>
					<Col xs={6} md={6}>
						<FragranceRatioInput value={this.props.fragrance_ratio} />
					</Col>
				</Row>
			</Grid>
		);
	}	
}

class SoapCalcLipidSelection extends Component {
	render() {
		return (
			<div className="container">
				<div className="row">
					<div className="col-md"><LipidLookup /></div>
    			</div>
  				<div className="row">
					<div className="col-md"><LipidTable lipids={this.props.lipids} /></div>
    			</div>
			</div>
		);
	}	
}

class LipidLookup extends Component {

	handleChange(e) {

	}

	render() {
		return (
			<FormGroup
				controlId="lipidLookup">
				<InputGroup>
					<FormControl
						type="text"
						bsSize="sm"
						placeholder="Find lipid..."
			 			onChange={this.handleChange}
						defaultValue={this.props.value} />
				</InputGroup>
				<FormControl.Feedback />
			</FormGroup>
		);
	}
}

class LipidTable extends Component {
	render() {
		const rows = [];

		this.props.lipids.forEach((lipid, i) => {
			rows.push(
				<LipidTableRow
				        key={i}
				        num={i+1}
						name={lipid.name}
						sap={lipid.sap}
						percentage={lipid.percentage}
						weight={lipid.weight} />
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
				<tbody>{rows}</tbody>
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
      			<td>{this.props.percentage}</td>
      			<td>{this.props.weight}</td>
				<td>delete</td>
    		</tr>
		);
	}
}

class Units extends Component {
	constructor(props) {
		super();
		this.state = {
      		activeButton: props.value,
    	};
	}

	onChange(e) {
		this.setState({
			activeButton: e,
		});
		//alert('Units ' + this.state.activeButton);
	}

	render() {
		return (
			<FormGroup controlId="units">
				<ControlLabel>Units</ControlLabel>
				<ButtonToolbar>
					<ToggleButtonGroup
						name="units"
						type="radio"
						onChange={this.onChange}
						defaultValue={this.state.activeButton}>
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
		super();
		this.state = {
      		activeButton: props.value,
    	};
	}
	
	render() {
		return (
			<FormGroup controlId="lyeType">
   				<ControlLabel>Lye Type</ControlLabel>
				<ButtonToolbar>
					<ToggleButtonGroup
						name="lye_type"
						type="radio"
						defaultValue={this.state.activeButton}>
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

	getValidationState () {
		/*
    	const length = this.state.value.length;
    	if (length > 10) return 'success';
    	else if (length > 5) return 'warning';
    	else if (length > 0) return 'error';
		*/
	}

	handleChange (event) {
		this.setState({ value: event.target.value });
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
				<PercentageInput name="water_lipid_ratio" value={this.props.value} />
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
				<PercentageInput name="super_fat_percentage" value={this.props.value} />
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
				<PercentageInput name="fragrance_ratio" value={this.props.value} />
				<FormControl.Feedback />
            	<HelpBlock>Validation is based on string length.</HelpBlock>
			</FormGroup>
		)
	}
}

class PercentageInput extends Component {

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

	handleChange (event) {
		this.setState({ value: event.target.value/100 });
	}

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

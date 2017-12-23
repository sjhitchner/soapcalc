import React, { Component } from 'react';
import { 
	Button, 
	ButtonGroup, 
	ButtonToolbar,
	/*
	FormGroup,
	ControlLabel,
	FormControl,
	HelpBlock,
	*/
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
			<div className="container">
				<div className="row">
					<SoapCalcParameters 
						lye_type={this.props.recipe.lye_type}
						lipid_weight={this.props.recipe.lipid_weight}
						water_to_lipid_ratio={this.props.recipe.water_to_lipid_ratio}
						super_fat_percentage={this.props.recipe.super_fat_percentage}
					/>
				</div>
				<div className="row">
					<SoapCalcLipidSelection lipids={this.props.recipe.lipids}/>
				</div>
			</div>
		);
	}
}

class SoapCalcParameters extends Component {
	render() {
		return (
			<table className="table">
				<tbody>
					<tr>
						<td>Lye Type</td>
						<td><LyeType type={this.props.lye_type}/></td>
						<td>Superfat</td>
						<td>{this.props.super_fat_percentage}</td>
					</tr>
					<tr>
						<td>Lipid Weight</td>
						<td>{this.props.lipid_weight}</td>
						<td>Water:Lipid Ratio</td>
						<td>{this.props.water_to_lipid_ratio}</td>
					</tr>
				</tbody>
			</table>
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
	render() {
		return (
			<input type="text" />
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

class UnitsSelector extends Component {
	render() {
		return (
			<ButtonToolbar>
				<ButtonGroup bsSize="large">
					<Button>Ounces</Button>
					<Button>Grams</Button>
				</ButtonGroup>
			</ButtonToolbar>
		)
	}
}

class LyeType extends Component {
	render() {
		return (
			<ButtonToolbar>
				<ButtonGroup>
					<Button>NaOH</Button>
					<Button>KOH</Button>
				</ButtonGroup>
			</ButtonToolbar>
		)
	}
}

class LipidWeight extends Component {
	render() {
		return (
			<div>
				Lipid Weight
			</div>
		)
	}
}

export default SoapCalc;

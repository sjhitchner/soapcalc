import React, { Component } from 'react';
import {
  BrowserRouter as Router,
  Route
} from 'react-router-dom';
//import logo from './logo.svg';
import '../../public/App.css';
import 'bootstrap/dist/css/bootstrap.min.css';
import Calculator from './calculator/CalculatorComponent';

class App extends Component {
  render() {
    return (
      <Router>
        <div>
          <Route exact path='/' component={Calculator} />
          <Route path='/home'   component={Calculator} />
          <Route path='/calc'   component={Calculator} />
        </div>
      </Router>
    );
  }
}

export default App;
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



import React, { Component } from 'react';
import { 
	Button,
	ButtonToolbar,
	ControlLabel,
	FormControl,
	FormGroup,
	HelpBlock,
	InputGroup,
	ToggleButton, 
	ToggleButtonGroup, 
} from 'react-bootstrap';

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
					defaultValue={this.props.value} />
				<InputGroup.Addon>{this.props.units}</InputGroup.Addon>
			</InputGroup>
		)
	}
}

class PercentageInput extends Component {
  
	handleChange = (e) => {
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


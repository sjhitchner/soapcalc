import React, { Component } from 'react';
import { 
	ControlLabel,
	FormControl,
	FormGroup,
	HelpBlock,
} from 'react-bootstrap';
import UnitsInput from './UnitsInputComponent';

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

export default LipidWeight;

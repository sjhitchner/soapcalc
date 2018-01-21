import React, { Component } from 'react';
import { 
	ControlLabel,
	FormControl,
	FormGroup,
	HelpBlock,
	InputGroup,
} from 'react-bootstrap';
import PercentageInput from './PercentageInputComponent';

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

export default SuperFatInput;

import React, { Component } from 'react';
import { 
	ControlLabel,
	FormControl,
	FormGroup,
	HelpBlock,
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
        <HelpBlock>The superfat percentage to calculate lye discount.</HelpBlock>
			</FormGroup>
		)
	}
}

export default SuperFatInput;

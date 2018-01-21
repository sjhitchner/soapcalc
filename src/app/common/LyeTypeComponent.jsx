import React, { Component } from 'react';
import { 
	ButtonToolbar,
	ControlLabel,
	FormControl,
	FormGroup,
	HelpBlock,
	InputGroup,
	ToggleButton, 
	ToggleButtonGroup, 
} from 'react-bootstrap';

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

export default LyeType;


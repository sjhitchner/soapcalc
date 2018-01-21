import React, { Component } from 'react';
import { 
	ButtonToolbar,
	ControlLabel,
	FormControl,
	FormGroup,
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
            <ToggleButton value={'lbs'}>Pounds</ToggleButton>
  					<ToggleButton value={'kg'}>Kilograms</ToggleButton>
					</ToggleButtonGroup>
				</ButtonToolbar>
				<FormControl.Feedback />
			</FormGroup>
		)
	}
}

export default Units;

import React, { Component } from 'react';
import { 
	ControlLabel,
	FormControl,
	FormGroup,
	HelpBlock,
} from 'react-bootstrap';
import PercentageInput from './PercentageInputComponent';

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
        <HelpBlock>Amount of fragrance with respect to lipids/fats.</HelpBlock>
			</FormGroup>
		)
	}
}

export default FragranceRatioInput;  

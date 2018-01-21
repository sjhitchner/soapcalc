import React, { Component } from 'react';
import { 
	ControlLabel,
	FormControl,
	FormGroup,
	HelpBlock,
} from 'react-bootstrap';
import PercentageInput from './PercentageInputComponent';

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

export default WaterLipidRatioInput;

import React, { Component } from 'react';
import { 
	FormControl,
	InputGroup,
} from 'react-bootstrap';


class PercentageInput extends Component {
  
	handleChange = (e) => {
    const value = e.target.value;
    if (value >= 0 && value <= 100) {
      this.props.onChange(value/100);
    }
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
					value={this.props.value * 100} />
				<InputGroup.Addon>%</InputGroup.Addon>
			</InputGroup>
		)
	}
}

export default PercentageInput;


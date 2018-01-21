import React, { Component } from 'react';
import { 
	FormControl,
	InputGroup,
} from 'react-bootstrap';

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
					value={this.props.value} />
				<InputGroup.Addon>{this.props.units}</InputGroup.Addon>
			</InputGroup>
		)
	}
}

export default UnitsInput;

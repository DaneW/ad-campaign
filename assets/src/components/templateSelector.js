import React, { Component } from 'react';
import PropTypes from 'prop-types';

class TemplateSelector extends Component {
  static propTypes = {
    templateId: PropTypes.string,
    templates: PropTypes.array,
    handleSelect: PropTypes.func, 
  }

  handleChange = event => {
    const { handleSelect } = this.props;
    handleSelect(event.target.value);
  };

  render() {
    const { templates, templateId } = this.props;
    if (!templates || !templates.length) return null;
    return (
      <div>
        <label htmlFor="template-selector">Template: </label>
        <select name="template-selector" value={templateId} onChange={this.handleChange}>
          <option disabled value="default"> Select a Template </option>
          {templates.map(({ id, title }) => <option key={id} value={id}>{title}</option>)}
        </select>
      </div>
    );
  }
}

export default TemplateSelector;

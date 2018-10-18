import React, { Component } from 'react';
import PropTypes from 'prop-types';

class AdEditor extends Component {
  static propTypes = {
    resources: PropTypes.array,
    titles: PropTypes.array,
    copies: PropTypes.array,
    images: PropTypes.array,
    updateValue: PropTypes.func,
  }

  constructor(props) {
    super(props);
    this.state = {
      valueModifier: -1,
    };
  }

  handleChange = (event, obj, idx) => {
    const { updateValue, titles, copies, images } = this.props;
    let update;
    let property;
    switch(obj) {
      case "title":
        update = [ ...titles ];
        property = 'adTitle';
        break;
      case "copy":
        update = [ ...copies ];
        property = 'adCopy';
        break;
      case "image":
        update = [ ...images ];
        property = 'adImage';
        break;
      default:
        break;
    }
    if (update && property) {
      const value = event.target ? event.target.value : event;
      update[idx] = value;
      updateValue(property, update);
    }
  };

  changeValueModifier = (event, idx) => {
    const { resources } = this.props;
    const i = event.target.value;
    const valueModifier = i;
    this.setState({ valueModifier });
    if (i < resources.length) {
      this.handleChange(resources[i].Name, "title", idx);
      this.handleChange(resources[i].Description, "copy", idx);
      this.handleChange(resources[i].Image, "image", idx);
    }
  };

  renderEditors = () => {
    const { valueModifier } = this.state;
    const { titles, copies, images, resources } = this.props;
    return titles.map((_, idx) => (
      <div key={idx}>
        <hr/>
        <label>Title: </label>
        <input 
          disabled={valueModifier[idx] !== resources.length + ""} 
          value={titles[idx]} 
          onChange={e => this.handleChange(e, "title", idx)} 
        />
        <br/>
        <label>Copy: </label>
        <input 
          disabled={valueModifier[idx] !== resources.length + ""} 
          value={copies[idx]} 
          onChange={e => this.handleChange(e, "copy", idx)} 
        /><br/>
        <label>Image: </label>
        <input 
          disabled={valueModifier[idx] !== resources.length + ""} 
          value={images[idx]} 
          onChange={e => this.handleChange(e, "image", idx)} 
        />
      </div>
    ));
  };

  render() {
    return (
      <div>
        {this.renderEditors()}
      </div>
    );
  }

}

export default AdEditor;
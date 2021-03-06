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
      valueModifier: [-1, -1, -1],
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
        return;
    }
    if (update && property) {
      const value = event.target ? event.target.value : event;
      update[idx] = value;
      updateValue(property, update);
    }
  };

  changeValueModifier = (event, idx) => {
    const { resources } = this.props;
    const { valueModifier: currModifier } = this.state;
    const index = event.target.value;
    const valueModifier = [ ...currModifier ];
    // Updating the current index
    valueModifier[idx] = index;
    this.setState({ valueModifier });
    if (index < resources.length) {
      this.handleChange(resources[index].Name, "title", idx);
      this.handleChange(resources[index].Description, "copy", idx);
      this.handleChange(resources[index].Image, "image", idx);
    }
  };

  renderList = idx => {
    const { valueModifier } = this.state;
    const { resources } = this.props;
    return (
      <select onChange={e => this.changeValueModifier(e, idx)} value={valueModifier[idx]}>
        <option disabled value={-1}>Default</option>
        {resources.map(({ Name }, i) => <option key={i} value={i}>{Name}</option>)}
        <option value={resources.length}>Custom</option>
      </select>
    );
  }

  renderEditors = () => {
    const { valueModifier } = this.state;
    const { titles, copies, images, resources } = this.props;
    return titles.map((_, idx) => (
      <div key={idx}>
        <hr/>
        {this.renderList(idx)}<br/>
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
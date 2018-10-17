import React, { Component } from 'react';
import PropTypes from 'prop-types';

class CampaignCreator extends Component {
  static propTypes = {
    template: PropTypes.object,
    saveCampaign: PropTypes.func,
    resources: PropTypes.array,
  }

  constructor(props) {
    super(props);
    this.state = {
      title: null,
      adTitle: [],
      adCopy: [],
      adImage: [],
      objective: null,
    };
  }

  componentDidUpdate(prevProps) {
    const { template } = this.props;
    if (template && template !== prevProps.template) {
      const { template: { ads, compainObjective: objective, title } } = this.props;
      const adTitle = ads.reduce((acc, curr) => [...acc, curr.title], []);
      const adCopy = ads.reduce((acc, curr) => [...acc, curr.copy], []);
      const adImage = ads.reduce((acc, curr) => [...acc, curr.image], []);
      this.setState({
        title,
        adTitle,
        adCopy,
        adImage,
        objective,
      });
    }
  }

  updateObjective = event => {
    const objective = event.target.value;
    this.setState({ objective });
  }

  handleUpdate = (name, value) => {
    this.setState({
      [name]: value
    });
  }

  handleSave = event => {
    event.preventDefault();
    const { saveCampaign } = this.props;
    const { title, adTitle, adCopy, adImage, objective } = this.state;
    saveCampaign({
      template: title,
      adTitle,
      adCopy,
      adImage,
      objective,
    });
  };

  isValid = () => {
    const { adTitle, adCopy, adImage, objective } = this.state;
    if (!(adTitle && adCopy && adImage && objective)) return false;
    return true;
  }

  render() {
    const { template } = this.props;
    const { objective } = this.state;
    if (!template) return null;
    return (
      <form>
        <div onChange={this.updateObjective}>
          <label><input type="radio" value="LeadGeneration" checked={objective === "LeadGeneration"} /> LeadGeneration</label>
          <label><input type="radio" value="Conversions" checked={objective === "Conversions"} /> Conversions</label>
          <label><input type="radio" value="Impressions" checked={objective === "Impressions"} /> Impressions</label>
        </div>
        <input type="submit" value="Save" disabled={!this.isValid()} onClick={this.handleSave} />
      </form>
    );
  }
}

export default CampaignCreator;

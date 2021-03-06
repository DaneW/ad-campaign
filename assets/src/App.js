import React, { Component } from 'react';
import TemplateSelector from './components/templateSelector';
import CampaignCreator from './components/campaignCreator';

const API = {
  templates: "http://localhost:3000/api/v1/templates",
  resources: "http://localhost:3000/api/v1/products",
  save: "http://localhost:3000/api/v1/campaigns/create",
  publish: id => `http://localhost:3000/api/v1/campaigns/${id}/publish`,
}

class App extends Component {

  constructor(props) {
    super(props);
    this.state = {
      template: null,
      templates: [],
      resources: [],
      campaignId: null,
      error: null,
      loading: true,
    };
  }

  componentDidMount() {
    const templates = fetch(API.templates)
      .then(response => response.json())
      .then(templates => this.setState({ templates }))
      .catch(error => this.setState({ error }));
    const resources = fetch(API.resources)
      .then(response => response.json())
      .then(resources => this.setState({ resources }))
      .catch(error => this.setState({ error }));
    Promise.all([templates, resources]).then(() => {
      this.setState({ loading: false });
    });
  }

  reset = () => {
    this.setState({
      template: null,
      campaignId: null,
      error: null,
    });
  };

  selectTemplate = templateId => {
    const { templates } = this.state;
    const template = templates.find(t => t.id === templateId);
    this.setState({ template });
  }

  saveCampaign = campaign => {
    fetch(API.save , { method: "POST", body: JSON.stringify(campaign) })
      .then(response => response.json())
      .then(({ id }) => this.setState({ campaignId: id }))
      .catch(error => this.setState({ error }));
  };

  publishCampaign = () => {
    const { campaignId } = this.state;
    fetch(API.publish(campaignId) , { method: "PUT" })
      .then(response => response.json())
      .then(() => this.reset())
      .catch(error => this.setState({ error }));
  };

  render() {
    const { templates, resources, loading, template, campaignId } = this.state;
    if (loading) return <p>Loading...</p>;
    const templateId = template ? template.id : "default";
    return (
      <div>
        <a href="http://localhost:3000/api/v1/templates">Campaign Templates</a>{' '}
        <a href="http://localhost:3000/api/v1/campaigns/published">Published Campaigns</a><hr/>
        <TemplateSelector templateId={templateId} templates={templates} handleSelect={this.selectTemplate}/>
        <CampaignCreator resources={resources} template={template} saveCampaign={this.saveCampaign} />
        {template && <button disabled={!campaignId} onClick={this.publishCampaign}>Publish</button>}
      </div>
    );
  }
}

export default App;

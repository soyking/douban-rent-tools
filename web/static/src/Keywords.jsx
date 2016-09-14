import React from 'react'
import {Tag, Input} from 'antd'
import {Box} from 'react-polymer-layout'
import 'antd/dist/antd.min.css'

const Keywords = React.createClass({
  getInitialState() {
    return {
      count: 0,
      input: "",
      keywords: []
    }
  },

  removeKeyword(i) {
    const keywords = [...this.state.keywords].filter(keyword => keyword && (keyword.index !== i))
    this.setState({keywords:keywords})
  },

  addKeyword(e) {
    if (e.target.value === "") {return}
    let {keywords, count} = this.state
    count=count+1
    keywords.push({index:count, value:e.target.value})
    this.setState({count:count, keywords:keywords, input:""})
  },

  inputChange(e) {
    this.setState({input:e.target.value})
  },

  render() {
    const { keywords, input } = this.state
    return (
      <Box stretch wrap style={{borderStyle:"solid", borderWidth:1, borderColor:"green", borderRadius:5}}>
        {
          keywords.map(keyword =>
            <div key={keyword.index} style={{margin:6}}>
              <Tag key={keyword.index} closable afterClose={() => this.removeKeyword(keyword.index)}>
                {keyword.value}
              </Tag>
            </div>
          )
        }
        <div style={{padding:5}}>
          <Input value={input} style={{width:100}} onChange={this.inputChange} onPressEnter={this.addKeyword}/>
        </div>
      </Box>
    );
  },
});

module.exports = Keywords
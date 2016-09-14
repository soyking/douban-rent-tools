import React from 'react'
import ReactDOM from 'react-dom'
import {Item, Box} from 'react-polymer-layout'
import Keywords from './Keywords.jsx'
import {DatePicker} from 'antd';

const Root = React.createClass({
  render() {
    return (
      <Box centerJustified>
        <Box vertical style={{width:"40%"}}>
          <Box center centerJustified>
            <div style={{color:"green", paddingTop:100, paddingBottom:50, fontSize:30}}>DOUBAN RENT TOOLS</div>
          </Box>
          <Box vertical>
            <Box center style={{margin:"10px 10px 10px"}}>
              <div style={{color:"green", fontSize:25, marginRight:86}}>关键词</div>
              <Item flex><Keywords /></Item>
            </Box>
            <Box cneter style={{margin:"10px 10px 10px"}}>
              <div style={{color:"green", fontSize:25, marginRight:10}}>最后更新时间</div>
              <Item flex style={{marginTop:5}}>
                <DatePicker.RangePicker format="yyyy-MM-dd HH:mm:ss" showTime style={{width:"100%"}}/>
              </Item>
            </Box>
            <Box cneter style={{margin:"10px 10px 10px"}}>
              <div style={{color:"green", fontSize:25, marginRight:10}}>最后回复时间</div>
              <Item flex style={{marginTop:5}}>
                <DatePicker.RangePicker format="yyyy-MM-dd HH:mm:ss" showTime style={{width:"100%"}}/>
              </Item>
            </Box>
          </Box>
        </Box>
      </Box>
    );
  },
});

ReactDOM.render(
  <Root />,
  document.querySelector('.root'))
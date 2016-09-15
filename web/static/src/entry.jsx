import React from 'react'
import ReactDOM from 'react-dom'
import {Item, Box} from 'react-polymer-layout'
import Keywords from './Keywords.jsx'
import {DatePicker, Button, Table, message} from 'antd'
import 'antd/dist/antd.min.css'
import './css/patch.css'

const Root = React.createClass({
  getInitialState() {
    return {
      keywords: [],
      fromUpdateTime: null,
      toUpdateTime: null,
      fromLastReplyTime: null,
      toLastReplyTime: null,
      data: [],
    }
  },

  updateKeywords(keywords){
    this.setState({keywords:keywords})
  },

  updateUpdateTime(value){
    if (value.length!=2) {return}
    this.setState(
      {
        fromUpdateTime: parseInt(new Date(value[0]).getTime()/1000),
        toUpdateTime: parseInt(new Date(value[1]).getTime()/1000)
      }
    )
  },

  updateLastReplyTime(value){
    if (value.length!=2) {return}
    this.setState(
      {
        fromLastReplyTime: parseInt(new Date(value[0]).getTime()/1000),
        toLastReplyTime: parseInt(new Date(value[1]).getTime()/1000)
      }
    )
  },

  formatDate(date){
    if (!date) {return null}
    return date.replace("T"," ").replace("+08:00","")
  },

  query(){
    let params = this.state
    let body = {
      "from_update_time": params.fromUpdateTime || 0,
      "to_update_time": params.toUpdateTime || 0,
      "from_last_reply_time": params.fromLastReplyTime || 0,
      "to_last_reply_time": params.toLastReplyTime ||0,
    }
    body.keywords = []
    params.keywords.forEach(k => {body.keywords.push(k.value)})

    let req = new XMLHttpRequest();
    req.open("POST", "/query", true)
    let that = this
    req.onreadystatechange = function(){
      if (req.readyState == 4) {
        if (req.status == 200){
          let resp = JSON.parse(req.responseText)
          if (resp.error) {
            message.error(resp.error)
          } else {
            let data = []
            resp.result.forEach(t => {
              let topic = {}
              topic.key = t._id
              topic.title = {title: t.title || "", url: t._id ||""},
              topic.author = {author: t.author || "", url: t.author_url ||""}
              topic.updateTime = that.formatDate(t.topic_content.update_time) || ""
              topic.lastReplyTime = that.formatDate(t.last_reply_time) || ""
              topic.withPic = t.topic_content.with_pic || false
              topic.reply = t.reply || 0
              topic.like = t.topic_content.like || 0
              data.push(topic)
            })
            that.setState({data:data})
          }
        } else {
          message.error(req.status)
        }
      }
    }
    req.send(JSON.stringify(body))
  },

  render() {
    const columns = [{
      title: '标题',
      dataIndex: 'title',
      key: 'title',
      width: 200,
      render(title) {
        return <a href={title.url} target="_blank">{title.title}</a>
      }
    }, {
      title: '用户',
      dataIndex: 'author',
      key: 'author',
      render(author) {
        return <a href={author.url} target="_blank">{author.author}</a>
      }
    }, {
      title: '更新时间',
      dataIndex: 'updateTime',
      key: 'updateTime'
    }, {
      title: '回复时间',
      dataIndex: 'lastReplyTime',
      key: 'lastReplyTime',
    }, {
      title: '有图',
      dataIndex: 'withPic',
      key: 'withPic',
      render: withPic => {return withPic?"是":"否" }
    }, {
      title: '回复数',
      dataIndex: 'reply',
      key: 'reply',
    }, {
      title: '喜欢数',
      dataIndex: 'like',
      key: 'like',
    }]
    const dataSource = this.state.data

    return (
      <Box centerJustified>
        <Box vertical style={{width:"40%"}}>
          <Box center centerJustified>
            <div style={{color:"green", paddingTop:100, paddingBottom:50, fontSize:30}}>DOUBAN RENT TOOLS</div>
          </Box>
          <Box vertical>
            <Box center style={{margin:"10px"}}>
              <div style={{color:"green", fontSize:25, marginRight:86}}>关键词</div>
              <Item flex><Keywords onChange={this.updateKeywords}/></Item>
            </Box>
            <Box cneter style={{margin:"10px"}}>
              <div style={{color:"green", fontSize:25, marginRight:10}}>最后更新时间</div>
              <Item flex style={{marginTop:5}}>
                <DatePicker.RangePicker format="yyyy-MM-dd HH:mm:ss" showTime style={{width:"100%"}} onChange={this.updateUpdateTime}/>
              </Item>
            </Box>
            <Box cneter style={{margin:"10px"}}>
              <div style={{color:"green", fontSize:25, marginRight:10}}>最后回复时间</div>
              <Item flex style={{marginTop:5}}>
                <DatePicker.RangePicker format="yyyy-MM-dd HH:mm:ss" showTime style={{width:"100%"}} onChange={this.updateLastReplyTime}/>
              </Item>
            </Box>
            <Box endJustified style={{margin:"10px"}}>
              <Button type="primary" icon="search" onClick={this.query}>搜索</Button>
            </Box>
          </Box>
          <Table dataSource={dataSource} columns={columns} />
        </Box>
      </Box>
    );
  },
});

ReactDOM.render(
  <Root />,
  document.querySelector('.root'))
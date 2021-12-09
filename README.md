# onebot-sdk
onebot-sdk
## 功能点
### 消息
- [x] `send_private_msg` 发送私聊消息
- [x] `send_group_msg` 发送群消息
- [x] `send_msg` 发送消息
- [x] `delete_msg` 撤回消息
- [x] `get_msg` 获取消息
- [x] `get_forward_msg` 获取合并转发消息

### 群组
- [x] `set_group_kick` 群组踢人
- [x] `set_group_ban` 群组单人禁言
- [ ] `set_group_anonymous_ban` 群组匿名用户禁言(`SetGroupAnonymousBan dev`)
- [ ] `set_group_whole_ban` 群组全员禁言(`SetGroupWholeBan dev`)
- [ ] `set_group_admin` 群组设置管理员(`SetGroupAdmin dev`)
- [ ] `set_group_anonymous` 群组匿名(`SetGroupAnonymous dev`)
- [ ] `set_group_card` 设置群名片（群备注）(`SetGroupCard dev`)
- [ ] `set_group_name` 设置群名(`SetGroupName dev`)
- [ ] `set_group_leave` 退出群组(`SetGroupLeave dev`)
- [ ] `set_group_special_title` 设置群组专属头衔(`SetGroupSpecialTitle dev`)
- [ ] `set_group_add_request` 处理加群请求／邀请(`SetGroupAddRequest dev`)
- [ ] `get_group_info` 获取群信息(`GetGroupInfo dev`)
- [ ] `get_group_list` 获取群列表(`GetGroupList dev`)
- [ ] `get_group_member_info` 获取群成员信息(`GetGroupMemberInfo dev`)
- [ ] `get_group_member_list` 获取群成员列表(`GetGroupMemberListInfo dev`)
- [ ] `get_group_honor_info` 获取群荣誉信息(`GetGroupHonorInfo dev`)

### 好友
- [x] `send_like` 发送好友赞
- [x] `set_friend_add_request` 处理加好友请求
- [x] `get_friend_list` 获取好友列表

### 账号
- [x] `get_login_info` 获取登录号信息
- [x] `get_stranger_info` 获取陌生人信息
- [ ] `get_cookies` 获取 Cookies (`GetCookies dev`)
- [ ] `get_csrf_token` 获取 CSRF Token (`GetCSRFToken dev`)
- [ ] `get_credentials` 获取 QQ 相关接口凭证(`GetCredentials dev`)
- [ ] `get_record` 获取语音(`GetRecord dev`)
- [ ] `get_image` 获取图片(`GetImage dev`)

### 系统
- [x] `get_status` 获取运行状态
- [x] `get_version_info` 获取版本信息
- [x] `set_restart` 重启 OneBot 实现
- [x] `clean_cache` 清理缓存
- [ ] `can_send_image` 检查是否可以发送图片(`CanSendImage dev`)
- [ ] `can_send_record` 检查是否可以发送语音(`CanSendRecord dev`)

package constant

const ConnectTimeout = 2.0
const TIMEOUT = 3.0

var AllowMethods = []string{"get", "delete", "head", "options", "patch", "post", "put"}

const ContentTypeForm = "application/x-www-form-urlencoded"
const ContentTypeJson = "application/json"
const ContentTypeMultiPart = "multipart/form-data"

const DataChannel = "IDG_CHANNELS"
const FromAppidKey = "from_appid"
const FromAppkeyKey = "from_appkey"
const FromChannelKey = "from_channel"

const ToAppidKey = "appid"
const ToAppkeyKey = "appkey"
const ToChannel = "channel"
const ToChannelAlias = "alias"
const AccountIdKey = "account_id"
const SubOrgKeyKey = "sub_org_key"
const UserInfoKey = "user_info"
const CallStackKey = "call_stack"
const DefaultChannelAlias = "default"
const ISS = "ItfarmPhpSdk"

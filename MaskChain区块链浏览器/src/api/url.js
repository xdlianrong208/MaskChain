console.log(process.env)

const baseJson = {
  jsonrpc: '2.0',
  method: '',
  params: [],
  id: 2
}

export function reviseParam (data = [], type) {
  baseJson.params = data
  switch (type) {
    case 0:
      baseJson.method = 'eth_getTotalTransactionsCount'
      break
    case 1:
      baseJson.method = 'eth_getBlockByNumber'
      break
    case 2:
      baseJson.method = 'eth_getBlockByHash'
      break
    case 3:
      baseJson.method = 'eth_getTransactionByHash'
      break
    case 4:
      baseJson.method = 'eth_getTransactionReceipt'
      break
    case 5:
      baseJson.method = 'eth_getCMState'
      break
    case 6:
      baseJson.method = 'txpool_status'
      break
    case 7:
      baseJson.method = 'eth_getTotalTransactions'
      break
    default:
      return baseJson
  }
  return baseJson
}

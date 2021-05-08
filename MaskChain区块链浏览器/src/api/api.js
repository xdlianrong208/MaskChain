import { post } from './http'
import { reviseParam } from './url'

export function getTransactionsCount (data) {
  // const d = ['0x0', true]
  const data2 = reviseParam(data, 0)
  return post({
    method: 'post',
    data: data2
  })
}

export function getBlockByNumber (number, bean = true) {
  // const d = ['0x0', true]
  number = '0x' + number
  const data = [number, bean]
  const data2 = reviseParam(data, 1)
  return post({
    method: 'post',
    data: data2
  })
}

export function getBlockByHash (hashData, bean = true) {
  // const d = ['0x0', true]
  const data = [hashData, bean]
  const data2 = reviseParam(data, 2)
  return post({
    method: 'post',
    data: data2
  })
}

export function getTbTransactionByPkHash (hashData) {
  // const d = ['0x0', true]
  const data = [hashData]
  const data2 = reviseParam(data, 3)
  return post({
    method: 'post',
    data: data2
  })
}

export function getTbTransactionReceiptByPkHash (hashData) {
  // const d = ['0x0', true]
  const data = [hashData]
  const data2 = reviseParam(data, 4)
  return post({
    method: 'post',
    data: data2
  })
}

export function getCMCount () {
  const data2 = reviseParam([], 5)
  return post({
    method: 'post',
    data: data2
  })
}

export function getTxpoolPending () {
  const data2 = reviseParam([], 6)
  return post({
    method: 'post',
    data: data2
  })
}

export function getTransactionsForFourteenDays2 () {
  const data2 = reviseParam([], 7)
  return post({
    method: 'post',
    data: data2
  })
}

export function getTbTransactionInfo (blockHeight, transactionHash) {
  // 输入只有区块高度时，将该区块的所有交易显示出来
  if (transactionHash) {
    return getTbTransactionByPkHash(transactionHash)
  } else if (blockHeight) {
    blockHeight = parseInt(blockHeight).toString(16)
    return getBlockByNumber(blockHeight)
  }
}

export async function getTransactionsForFourteenDays (blockNumber) {
  const OneDayTimes = 24 * 60 * 60
  const NowTime = new Date()
  const TransactionsList = new Array(14)
  let ADayTransactions = []
  let lastBetweenDay = 0
  let nowBetweenDay = 0
  let NowDay = Math.floor(NowTime.getTime() / 1000)
  NowDay = NowDay - 60 * 60 * NowTime.getHours() - 60 * NowTime.getMinutes() - NowTime.getSeconds()
  let i = 0
  while (true) {
    const x = await getBlockByNumber((blockNumber - i).toString(16))
      .then((res) => {
        const time = parseInt(res.data.result.timestamp, 16)
        const t = new Date(time * 1000)
        // 区块打包的时间(天)
        const blockDay = time - 60 * 60 * t.getHours() - 60 * t.getMinutes() - t.getSeconds()
        nowBetweenDay = (NowDay - blockDay) / OneDayTimes
        if (i === 0) {
          lastBetweenDay = nowBetweenDay
        }
        if (nowBetweenDay <= 14) {
          i += 1
          if (lastBetweenDay === nowBetweenDay) {
            ADayTransactions = ADayTransactions.concat(res.data.result.transactions)
          } else {
            TransactionsList[lastBetweenDay] = ADayTransactions
            lastBetweenDay = nowBetweenDay
            ADayTransactions = []
            ADayTransactions = ADayTransactions.concat(res.data.result.transactions)
          }
          return true
        } else {
          TransactionsList[lastBetweenDay] = ADayTransactions
          ADayTransactions = []
          return false
        }
      })
    if (!x) {
      break
    }
  }
  return {
    TransactionsList: TransactionsList,
    NowDay: NowDay
  }
}

// const registerKey = (key, url) => {
//     KEYMAP.set(key, url);
// }

var keySeq = []
var nodeSeq = [KEYMAP]
var curr = nodeSeq[0]

const handleKey = (e) => {
    console.log(keySeq)
    // console.log(nodeSeq)
    // console.log(curr)
    var key = e.key.toLowerCase()
    if (key === 'escape') {
        back()
        return
    }
    if ('children' in curr && key in curr.children) {
        step(key)
        // console.log(keySeq)
        // node = node.children.get(key);
        // console.log(nodeSeq)
        // console.log(sequence);
        // if (e.shiftKey) {
        //     window.open(node.get(key)?.link, '_blank');
        // } else {
        //     window.open(node.get(key)?.link, '_self');
        // }
    }
}

document.onkeydown = function(e) {
    handleKey(e)
}

const go = (e, link) => {
    if (e.shiftKey) {
        window.open(link, '_blank')
    } else {
        window.open(link, '_self')
    }
}

const step = (key) => {
    keySeq.push(key)
    curr = curr.children[key]
    nodeSeq.push(curr)
    update()
}

const back = () => {
    if (keySeq.length === 0) {
        return
    }
    keySeq.pop()
    nodeSeq.pop()
    curr = nodeSeq[nodeSeq.length - 1]
    update()
}

const reset = () => {
    keySeq.splice(0)
    nodeSeq.splice(1)
    curr = nodeSeq[0]
    update()
}

const update = () => {
    console.log('updating')
    if (keySeq.length === 0) {
        document.getElementById('seqText').innerText = '...'
        return
    }
    document.getElementById('seqText').innerText = keySeq.join(' ')
}

// for (const item of CONFIG.firstlistsContainer) {
//     for (const link of item.links) {
// if (link.key) {
//     registerKey(link.key, link.link);
//     }
// }
// }

// for (const item of CONFIG.secondListsContainer) {
//     for (const link of item.links) {
// if (link.key) {
//     registerKey(link.key, link.link);
//     }
// }
// }

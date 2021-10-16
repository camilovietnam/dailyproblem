<table>
    <thead>
        <th class="th-root">
            Root
        </th>
    </thead>
</table>

<div>
    <button class="add-row" on:click={addRow}>Add row</button>    
    <button on:click={calculateSum}>Calculate sum</button>    
</div>

<div class='div-result'>
    
</div>

<style>
    table {
        width: 100%;
        border: 1px solid black;
    }

    button {
        margin-top: 20px;
        padding: 20px;
        width: 100px;
        cursor: pointer;
	}
</style> 

<script lang="ts">
    interface Node {
        value: number,
        left: Node,
        right: Node,
    }

    // Depth is the number of cells to add on next iteration
    // it starts with 1 cell for the first level added. 
    let cellsToAdd: number = 1; 

    function addRow () {
        if (cellsToAdd > 16) {
            // The HTML will break if we add too many cells, so let's 
            // prevent the user from doing so. 
            alert('Let\'s not add too many levels');
            j('.add-row').attr('disabled', 'disabled')

            return
        }

        const row: any = j(`<tr></tr>`)

        for (let i = 0; i < cellsToAdd; i++) {
            j(row).append(`<td class="border-red"><input style="width: 50px" type="text" /></td>`) 
        }

        cellsToAdd *= 2

        j('table').append(row);

        updateColSpans()
    }

    function updateColSpans() {
        let colspan: number = cellsToAdd

        // We update from the bottom row to the top one
        j('table tr').each((i, row) => {
            j('td', row).attr('colspan', colspan)
            colspan /= 2
        })

        // update colspan of header to cover all the width
        j('.th-root').attr('colspan', cellsToAdd)
    }

    function calculateSum (){
        const tree: Node = buildTree() 
        const paths: number[][] = getPaths(tree)
        
        let minSum: number
        let minPath: number[]

        paths.forEach(path => {
            const sum: number = path.reduce((prev: number, current: number) => {
                return prev + current
            })

            if (!minSum || sum < minSum) {
                minSum = sum
                minPath = path
            }
        })

        j('.div-result').html(`
            The min cost: ${minSum}
            The branch is: [${minPath}]
        `)
    }

    function buildTree () {
        const rows: any[] = j('tr').get()
        let parentNodes: Node[] = []
        const root: Node = {
            value: null,
            left: null,
            right: null,
        } 

        for (let row in rows) {
            const rowIndex: number = parseInt(row)
            const cols: any[] = j('input', rows[rowIndex]).get()
            const newParentNodes: Node[] = []
            
            if (rowIndex == 0) {
                // For the root node
                const value = cols[0].value
                root.value = parseInt(value);
                newParentNodes.push(root)
            } else {
                // For the leaf nodes
                for (let col in cols) {
                    const colIndex: number = parseInt(col)
                    const node: Node = {
                        left: null,
                        right: null,
                        value: parseInt(cols[col].value),
                    }

                    if (colIndex%2 == 0) {
                        addToParentAsLeft(node, col, parentNodes)
                    } else {
                        addToParentAsRight(node, col, parentNodes)
                    }

                    newParentNodes.push(node)
                }                
            }            

            parentNodes = newParentNodes
        }

        return root
    }

    function addToParentAsLeft (node, nodeIndex, parents) {
        const parentIndex = Math.floor(nodeIndex/2)
        parents[parentIndex].left = node
    }

    function addToParentAsRight (node, nodeIndex, parents) {
        const parentIndex = Math.floor(nodeIndex/2);
        parents[parentIndex].right = node
    }

    function getPaths(node: any) {
        if (isEmpty(node)) {
            return []
        }

        if (hasNoChildren(node)) {
            return [[node.value]];
        }

        let leftPaths = getPaths(node.left)
        let rightPaths = getPaths(node.right)        

        leftPaths = addToBranches(node.value, leftPaths)
        rightPaths = addToBranches(node.value, rightPaths)

        return [...leftPaths, ...rightPaths]
    }

    function addToBranches(value: number, branches: number[][]) {
        return branches.map(path => {
            return [value, ...path]
        })
    }

    function hasNoChildren(node: Node) {
        return (!node.left && !node.right) || (!node.left.value && !node.right.value)
    }

    function isEmpty(node: Node) {
        return !node || !node.value
    }
</script>
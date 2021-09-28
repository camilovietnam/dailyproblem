<script lang="ts">
	import Problem from './Problem.svelte'
	let candidates : number[][] = [
		[-9,-8,-7],
		[-3,-2,-1],
		[0,1,2],
		[3,4,5],
	]

	let list : number [] = []
	let squaredList : number[] = []

	function addNumber(e) {
		list = [...list, parseInt(e.target.value)]
	}

	function resetList() {
		list = []
		squaredList = []
	}

	function squareList() {
		squaredList = list.map((num) => {
			console.log("I will square: ", num)
			return num**2
		})
		squaredList = squaredList.sort((a,b) => (a-b))
	}

</script>

<main>
	<h1>Sorted Squares</h1>

	<Problem />

	{#each candidates as row}
		{#each row as number}
			<button class="button-number" on:click="{addNumber}" value={number}>{number}</button>
		{/each}
		<br />
	{/each}
	<hr />
	List:
	<br />
	{list.join(',')}
	<br />
	<br />
	<button on:click={squareList}>Square and sort</button>
	<button on:click={resetList}>Reset</button>

	<p>
		Squared list: {squaredList}
	</p>
</main>

<style>
	.button-number {
		padding: 20px;
		width: 100px;
		cursor: pointer;
	}

	main {
		text-align: center;
		padding: 1em;
		max-width: 240px;
		margin: 0 auto;
	}

	h1 {
		color: #ff3e00;
		text-transform: uppercase;
		font-size: 4em;
		font-weight: 100;
		margin: 0 auto;
	}

	@media (min-width: 640px) {
		main {
			max-width: none;
		}
	}
</style>
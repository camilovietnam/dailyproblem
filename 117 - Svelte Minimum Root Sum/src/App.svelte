<script lang="ts">
	const jqueryScript = document.createElement('script');
	jqueryScript.setAttribute('src', 'jquery.min.js')
	jqueryScript.onload = function () {
		window.j = jQuery;
	}

	document.querySelector('body').appendChild(jqueryScript)

	import Problem from './Problem.svelte'
	import TreeForm from './TreeForm.svelte' 
</script>

<main>
	<h1>Minimum Root Sum</h1>

	<Problem />
	<TreeForm />

</main>

<style>
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
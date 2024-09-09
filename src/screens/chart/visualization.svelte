<script lang="ts">
  import { getWeights } from "../../api/api";
  import Chart from "./chart.svelte";
  import ChartForm from "./chart-form.svelte";
  import type { Range } from "./types";
  import { readRange, storeRange } from "./persist-range";
  import { switchScreen } from "../../store/screen";
  import Page from "../../ui/page.svelte";
  import { i18n } from "../../store/language";

  const handleShowAddWeight = () => switchScreen("addWeight");

  let selectedRange: Range = readRange() ?? "14-days";

  const onSelectRange = (value: Range) => {
    selectedRange = value;
    storeRange(value);
    weightsRequest = getWeights(selectedRange);
  };

  let weightsRequest = getWeights(selectedRange);
</script>

<Page column={true}>
  {#await weightsRequest}
    {$i18n("chart.loading")}
  {:then result}
    <Chart weights={result.isSuccess ? result.data : []} />
    <ChartForm {handleShowAddWeight} {selectedRange} {onSelectRange} />
  {:catch error}
    {$i18n("chart.errorOccurred")}: {error.message}
  {/await}
</Page>

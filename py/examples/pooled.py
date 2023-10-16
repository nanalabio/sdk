import nanala

experiment = nanala.new_experiment()

partA = ["1","2","3"]
partB = ["I","II","III"]
barcodes = nanala.barcodes["setA"]

pool_clone = experiment.golden_gate_pooled(partA,partB, barcodes=barcodes) # Define what the pool object is
pool_transformation = experiment.transform_pool(plasmids=pool_clone, strain="Escherichia coli BL21(DE3)")
pool_nanopore_sequencing = experiment.nanopore(pool_transformation, primers=nanala.primers["M13"])
sorted_pool = experiment.facs(pool_clone, sort_by="fluorescence")
pool_illumina_sequencing = experiment.illumina(sorted_pool)
experiment.output({"nanopore_sequencing": pool_nanopore_sequencing, "illumina_sequencing": pool_illumina_sequencing})

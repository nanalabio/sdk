import nanala

# Our experiment will involve testing to see if a variety of different T7RNAp will express
# GFP. We will synthesize a variety of different T7RNAp,

# First, let's clone our protein into an expression vector
T7RNAp = "MNTINIAKNDFSDIELAAIPFNTLADHYGERLAREQLALEHESYEMGEARFRKMFERQLKAGEVADNAAAKPLITTLLPKMIARINDWFEEVKAKRGKRPTAFQFLQEIKPEAVAYITIKTTLACLTSADNTTVQAVASAIGRAIEDEARFGRIRDLEAKHFKKNVEEQLNKRVGHVYKKAFMQVVEADMLSKGLLGGEAWSSWHKEDSIHVGVRCIEMLIESTGMVSLHRQNAGVVGQDSETIELAPEYAEAIATRAGALAGISPMFQPCVVPPKPWTGITGGGYWANGRRPLALVRTHSKKALMRYEDVYMPEVYKAINIAQNTAWKINKKVLAVANVITKWKHCPVEDIPAIEREELPMKPEDIDMNPEALTAWKRAAAAVYRKDKARKSRRISLEFMLEQANKFANHKAIWFPYNMDWRGRVYAVSMFNPQGNDMTKGLLTLAKGKPIGKEGYYWLKIHGANCAGVDKVPFPERIKFIEENHENIMACAKSPLENTWWAEQDSPFCFLAFCFEYAGVQHHGLSYNCSLPLAFDGSCSGIQHFSAMLRDEVGGRAVNLLPSETVQDIYGIVAKKVNEILQADAINGTDNEVVTVTDENTGEISEKVKLGTKALAGQWLAYGVTRSVTKRSVMTLAYGSKEFGFRQQVLEDTIQPAIDSGKGLMFTQPNQAAGYMAKLIWESVSVTVVAAVEAMNWLKSAAKLLAAEVKDKKTGEILRKRCAVHWVTPDGFPVWQEYKKPIQTRLNLMFLGQFRLQPTINTNKDSEIDAHKQESGIAPNFVHSQDGSHLRKTVVWAHEKYGIESFALIHDSFGTIPADAANLFKAVRETMVDTYESCDVLADFYDQFADQLHESQLDKMPALPAKGNLNLRDILESDFAFA"

T7RNAPS = [T7RNAp,] # Add more here later
for T7RNAp in T7RNAPS:
    t7_gene = nanala.synthesize_cds(amino_acids=T7RNAP, host="Escherichia coli")
    
    # Next, let's simulate a GoldenGate assembly using our synthesized T7, a few selected parts,
    # and a GFP test construct. We're going to pretend that GFP_TEST is a real test construct
    # that will be placed at the end of our GoldenGate.
    #
    # All the other parts are prebuilt parts in Nanala's cache, ready for use!
    vector = nanala.golden_gate(promoter="araBAD", rbs="BBa_B0034", cds=t7_gene, terminator="tonB", suffix=GFP_TEST, resistance="kanamycin", origin="p15a")
    
    # Next, let's transform our cloned vector into E.coli turbos. Note: you can
    # directly use the vector, but there aren't any guarantees about the host.
    # We may clone into either Vibrio natriegens or a wide array of different
    # Escherichia coli strains. This transformation makes it strain specific.
    expression_strain = nanala.transform(plasmid=vector, strain="Escherichia coli NEB turbo")
    
    # Finally, let's run a plate assay with the expression strain.
    expression_data = nanala.plate_assay(measure=["od600","gfp"], induce="arabinose", strain=expression_strain)
    
    # We can get the expression data from this id later!
    print(expression_data.id)
    
    # ... LATER ...
    data = nanala.get_data(expression_data.id)
    print(data) # TSV file of biological data